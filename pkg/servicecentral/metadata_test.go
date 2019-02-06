package servicecentral

import (
	"testing"

	creator_v1 "github.com/atlassian/voyager/pkg/apis/creator/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetsAndGetsMetadataCorrectly(t *testing.T) {
	t.Parallel()

	serviceData := &ServiceDataWrite{}

	pdMetadata := creator_v1.PagerDutyMetadata{
		Staging: creator_v1.PagerDutyEnvMetadata{
			Main: creator_v1.PagerDutyServiceMetadata{
				ServiceID: "some-service",
				PolicyID:  "some-policy",
				Integrations: creator_v1.PagerDutyIntegrations{
					Generic: creator_v1.PagerDutyIntegrationMetadata{
						IntegrationID:  "some-id",
						IntegrationKey: "some-key",
					},
				},
			},
		},
	}

	err := SetPagerDutyMetadata(serviceData, &pdMetadata)
	require.NoError(t, err)

	actual, err := GetPagerDutyMetadata(serviceData)

	require.NoError(t, err)
	assert.Equal(t, &pdMetadata, actual)
}

func TestGetMetadataReturnsErrorWhenMiscHasInvalidJSON(t *testing.T) {
	t.Parallel()

	// we have to know how it's stored in the misc field unfortunately...
	serviceData := &ServiceDataRead{
		ServiceDataWrite: ServiceDataWrite{
			Misc: []miscData{
				{
					Key:   PagerDutyMetadataKey,
					Value: "{ \"foo\": notvalid }", // not valid json
				},
			},
		},
	}

	_, err := GetPagerDutyMetadata(&serviceData.ServiceDataWrite)

	require.Error(t, err)
}

func TestGetMetadataReturnsNilWhenNoPreviousData(t *testing.T) {
	t.Parallel()

	actual, err := GetPagerDutyMetadata(&ServiceDataWrite{})

	require.NoError(t, err)
	assert.Nil(t, actual)
}

func TestSetsAndGetsBuildsCorrectly(t *testing.T) {
	t.Parallel()

	serviceData := &ServiceDataRead{}

	buildMetadata := creator_v1.BambooMetadata{
		Builds: []creator_v1.BambooPlanRef{
			{
				Server: "sox-bamboo",
				Plan:   "VYGR-VYGR",
			},
		},
		Deployments: []creator_v1.BambooPlanRef{
			{
				Server: "staging-bamboo",
				Plan:   "VYGR-VYGR",
			},
		},
	}

	err := SetBambooMetadata(&serviceData.ServiceDataWrite, &buildMetadata)
	require.NoError(t, err)

	actual, err := GetBambooMetadata(&serviceData.ServiceDataWrite)

	require.NoError(t, err)
	assert.Equal(t, &buildMetadata, actual)
}

func TestGetBuildsReturnsErrorWhenMiscHasInvalidJSON(t *testing.T) {
	t.Parallel()

	// we have to know how it's stored in the misc field unfortunately...
	serviceData := &ServiceDataRead{
		ServiceDataWrite: ServiceDataWrite{
			Misc: []miscData{
				{
					Key:   BambooMetadataKey,
					Value: "{ \"foo\": notvalid }", // not valid json
				},
			},
		},
	}

	_, err := GetBambooMetadata(&serviceData.ServiceDataWrite)

	require.Error(t, err)
}

func TestGetBuildsReturnsNilWhenNoPreviousData(t *testing.T) {
	t.Parallel()

	actual, err := GetBambooMetadata(&ServiceDataWrite{})

	require.NoError(t, err)
	assert.Nil(t, actual)
}
