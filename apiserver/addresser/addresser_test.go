// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package addresser_test

import (
	"errors"

	gc "gopkg.in/check.v1"

	jc "github.com/juju/testing/checkers"

	"github.com/juju/juju/apiserver/addresser"
	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/params"
	apiservertesting "github.com/juju/juju/apiserver/testing"
	"github.com/juju/juju/cmd/envcmd"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/environs/configstore"
	"github.com/juju/juju/feature"
	"github.com/juju/juju/provider/dummy"
	"github.com/juju/juju/state"
	statetesting "github.com/juju/juju/state/testing"
	coretesting "github.com/juju/juju/testing"
)

type AddresserSuite struct {
	coretesting.BaseSuite

	st         *mockState
	api        *addresser.AddresserAPI
	authoriser apiservertesting.FakeAuthorizer
	resources  *common.Resources
}

var _ = gc.Suite(&AddresserSuite{})

func (s *AddresserSuite) SetUpTest(c *gc.C) {
	s.BaseSuite.SetUpTest(c)
	s.SetFeatureFlags(feature.AddressAllocation)

	s.authoriser = apiservertesting.FakeAuthorizer{
		EnvironManager: true,
	}
	s.resources = common.NewResources()
	s.AddCleanup(func(*gc.C) { s.resources.StopAll() })

	s.st = newMockState()
	addresser.PatchState(s, s.st)

	var err error
	s.api, err = addresser.NewAddresserAPI(nil, s.resources, s.authoriser)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *AddresserSuite) TearDownTest(c *gc.C) {
	dummy.Reset()
	s.BaseSuite.TearDownTest(c)
}

func (s *AddresserSuite) TestCleanupIPAddressesSuccess(c *gc.C) {
	config := testingEnvConfig(c)
	s.st.setConfig(c, config)

	dead, err := s.st.DeadIPAddresses()
	c.Assert(err, gc.IsNil)
	c.Assert(dead, gc.HasLen, 2)

	apiErr := s.api.CleanupIPAddresses()
	c.Assert(apiErr.Error, gc.IsNil)

	dead, err = s.st.DeadIPAddresses()
	c.Assert(err, gc.IsNil)
	c.Assert(dead, gc.HasLen, 0)
}

func (s *AddresserSuite) TestCleanupIPAddressesFailure(c *gc.C) {
	config := testingEnvConfig(c)
	s.st.setConfig(c, config)

	dead, err := s.st.DeadIPAddresses()
	c.Assert(err, gc.IsNil)
	c.Assert(dead, gc.HasLen, 2)

	s.st.stub.SetErrors(errors.New("ouch"))

	// First action is getting the environment configuration,
	// so the injected error is returned here.
	apiErr := s.api.CleanupIPAddresses()
	c.Assert(apiErr.Error, gc.ErrorMatches, "getting environment config: ouch")

	// Still has two dead addresses.
	dead, err = s.st.DeadIPAddresses()
	c.Assert(err, gc.IsNil)
	c.Assert(dead, gc.HasLen, 2)
}

func (s *AddresserSuite) TestWatchIPAddresses(c *gc.C) {
	c.Assert(s.resources.Count(), gc.Equals, 0)

	s.st.addIPAddressWatcher("0.1.2.3", "0.1.2.4", "0.1.2.7")

	result, err := s.api.WatchIPAddresses()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, gc.DeepEquals, params.EntityWatchResult{
		EntityWatcherId: "1",
		Changes: []string{
			"ipaddress-00000000-1111-2222-3333-0123456789ab",
			"ipaddress-00000000-1111-2222-4444-0123456789ab",
			"ipaddress-00000000-1111-2222-7777-0123456789ab",
		},
		Error: nil,
	})

	// Verify the resource was registered and stop when done.
	c.Assert(s.resources.Count(), gc.Equals, 1)
	resource := s.resources.Get("1")
	defer statetesting.AssertStop(c, resource)

	// Check that the Watch has consumed the initial event ("returned" in
	// the Watch call)
	wc := statetesting.NewStringsWatcherC(c, s.st, resource.(state.StringsWatcher))
	wc.AssertNoChange()
}

// testingEnvConfig prepares an environment configuration using
// the dummy provider.
func testingEnvConfig(c *gc.C) *config.Config {
	cfg, err := config.New(config.NoDefaults, dummy.SampleConfig())
	c.Assert(err, jc.ErrorIsNil)
	env, err := environs.Prepare(cfg, envcmd.BootstrapContext(coretesting.Context(c)), configstore.NewMem())
	c.Assert(err, jc.ErrorIsNil)
	return env.Config()
}
