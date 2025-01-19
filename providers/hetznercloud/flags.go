package hetznercloud

import (
	"os"

	"github.com/urfave/cli/v2"
)

const category = "Hetzner Cloud"

var ProviderFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "hetznercloud-api-token",
		Usage:    "hetzner cloud api token",
		EnvVars:  []string{"CROW_HETZNERCLOUD_API_TOKEN", "WOODPECKER_HETZNERCLOUD_API_TOKEN"},
		FilePath: os.Getenv("CROW_HETZNERCLOUD_API_TOKEN_FILE"),
		Category: category,
	},
	&cli.StringFlag{
		Name:     "hetznercloud-location",
		Value:    "nbg1",
		Usage:    "hetzner cloud location",
		EnvVars:  []string{"CROW_HETZNERCLOUD_LOCATION", "WOODPECKER_HETZNERCLOUD_LOCATION"},
		Category: category,
	},
	&cli.StringFlag{
		Name:     "hetznercloud-server-type",
		Value:    "cx11",
		Usage:    "hetzner cloud server type",
		EnvVars:  []string{"CROW_HETZNERCLOUD_SERVER_TYPE", "WOODPECKER_HETZNERCLOUD_SERVER_TYPE"},
		Category: category,
	},
	&cli.StringSliceFlag{
		Name:     "hetznercloud-ssh-keys",
		Usage:    "names of hetzner cloud ssh keys",
		EnvVars:  []string{"CROW_HETZNERCLOUD_SSH_KEYS", "WOODPECKER_HETZNERCLOUD_SSH_KEYS"},
		Category: category,
	},
	&cli.StringFlag{
		Name:     "hetznercloud-user-data",
		Usage:    "hetzner cloud userdata template",
		EnvVars:  []string{"CROW_HETZNERCLOUD_USERDATA", "WOODPECKER_HETZNERCLOUD_USERDATA"},
		FilePath: os.Getenv("CROW_HETZNERCLOUD_USERDATA_FILE"),
		Category: category,
	},
	&cli.StringFlag{
		Name:     "hetznercloud-image",
		Value:    "ubuntu-22.04",
		Usage:    "hetzner cloud image",
		EnvVars:  []string{"CROW_HETZNERCLOUD_IMAGE", "WOODPECKER_HETZNERCLOUD_IMAGE"},
		Category: category,
	},
	&cli.StringSliceFlag{
		Name:     "hetznercloud-labels",
		Usage:    "hetzner cloud server labels",
		EnvVars:  []string{"CROW_HETZNERCLOUD_LABELS", "WOODPECKER_HETZNERCLOUD_LABELS"},
		Category: category,
	},
	&cli.StringSliceFlag{
		Name:     "hetznercloud-firewalls",
		Usage:    "names of hetzner cloud firewalls",
		EnvVars:  []string{"CROW_HETZNERCLOUD_FIREWALLS", "WOODPECKER_HETZNERCLOUD_FIREWALLS"},
		Category: category,
	},
	&cli.StringSliceFlag{
		Name:     "hetznercloud-networks",
		Usage:    "names of hetzner cloud networks",
		EnvVars:  []string{"CROW_HETZNERCLOUD_NETWORKS", "WOODPECKER_HETZNERCLOUD_NETWORKS"},
		Category: category,
	},
	&cli.BoolFlag{
		Name:     "hetznercloud-public-ipv4-enable",
		Value:    true,
		Usage:    "enables public ipv4 network for agents",
		EnvVars:  []string{"CROW_HETZNERCLOUD_PUBLIC_IPV4_ENABLE", "WOODPECKER_HETZNERCLOUD_PUBLIC_IPV4_ENABLE"},
		Category: category,
	},
	&cli.BoolFlag{
		Name:     "hetznercloud-public-ipv6-enable",
		Value:    true,
		Usage:    "enables public ipv6 network for agents",
		EnvVars:  []string{"CROW_HETZNERCLOUD_PUBLIC_IPV6_ENABLE", "WOODPECKER_HETZNERCLOUD_PUBLIC_IPV6_ENABLE"},
		Category: category,
	},
}
