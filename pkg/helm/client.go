package helm

import (
	"google.golang.org/grpc"
	"k8s.io/helm/pkg/chartutil"
	rls "k8s.io/helm/pkg/proto/hapi/services"
	"os"
)

const (
	// $HELM_HOST envvar
	HelmHostEnvVar = "HELM_HOST"

	// $HELM_HOME envvar
	HelmHomeEnvVar = "HELM_HOME"

	// Default tiller server host address.
	DefaultHelmHost = ":44134"

	// Default $HELM_HOME envvar value
	DefaultHelmHome = "$HOME/.helm"
)

// Helm client manages client side of the helm-tiller protocol
type Client struct {
	opts options
}

func NewClient(opts ...Option) *Client {
	return new(Client).Init().Option(opts...)
}

// Configure the helm client with the provided options
func (h *Client) Option(opts ...Option) *Client {
	for _, opt := range opts {
		opt(&h.opts)
	}
	return h
}

// Initializes the helm client with default options
func (h *Client) Init() *Client {
	return h.Option(HelmHost(DefaultHelmHost)).
		Option(HelmHome(os.ExpandEnv(DefaultHelmHome)))
}

// ListReleases lists the current releases.
func (h *Client) ListReleases(opts ...ReleaseListOption) (*rls.ListReleasesResponse, error) {
	c, err := grpc.Dial(h.opts.host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return h.opts.rpcListReleases(rls.NewReleaseServiceClient(c), opts...)
}

// InstallRelease installs a new chart and returns the release response.
func (h *Client) InstallRelease(chStr string, opts ...InstallOption) (*rls.InstallReleaseResponse, error) {
	c, err := grpc.Dial(h.opts.host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer c.Close()

	chart, err := chartutil.Load(chStr)
	if err != nil {
		return nil, err
	}

	return h.opts.rpcInstallRelease(chart, rls.NewReleaseServiceClient(c), opts...)
}

// UninstallRelease uninstalls a named release and returns the response.
//
// Note: there aren't currently any supported DeleteOptions, but they are
// kept in the API signature as a placeholder for future additions.
func (h *Client) DeleteRelease(rlsName string, opts ...DeleteOption) (*rls.UninstallReleaseResponse, error) {
	c, err := grpc.Dial(h.opts.host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return h.opts.rpcDeleteRelease(rlsName, rls.NewReleaseServiceClient(c), opts...)
}

// UpdateRelease updates a release to a new/different chart.
//
// Note: there aren't currently any supported UpdateOptions, but they
// are kept in the API signature as a placeholder for future additions.
func (h *Client) UpdateRelease(rlsName string, opts ...UpdateOption) (*rls.UpdateReleaseResponse, error) {
	c, err := grpc.Dial(h.opts.host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return h.opts.rpcUpdateRelease(rlsName, rls.NewReleaseServiceClient(c), opts...)
}

// ReleaseStatus returns the given release's status.
//
// Note: there aren't currently any  supported StatusOptions,
// but they are kept in the API signature as a placeholder for future additions.
func (h *Client) ReleaseStatus(rlsName string, opts ...StatusOption) (*rls.GetReleaseStatusResponse, error) {
	c, err := grpc.Dial(h.opts.host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return h.opts.rpcGetReleaseStatus(rlsName, rls.NewReleaseServiceClient(c), opts...)
}

// ReleaseContent returns the configuration for a given release.
//
// Note: there aren't currently any supported ContentOptions, but
// they are kept in the API signature as a placeholder for future additions.
func (h *Client) ReleaseContent(rlsName string, opts ...ContentOption) (*rls.GetReleaseContentResponse, error) {
	c, err := grpc.Dial(h.opts.host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return h.opts.rpcGetReleaseContent(rlsName, rls.NewReleaseServiceClient(c), opts...)
}
