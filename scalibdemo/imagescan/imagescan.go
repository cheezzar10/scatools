package imagescan

import (
	"context"
	"fmt"
	scalibr "github.com/google/osv-scalibr"
	layerimg "github.com/google/osv-scalibr/artifact/image/layerscanning/image"
	"github.com/google/osv-scalibr/plugin/config"
	pluglist "github.com/google/osv-scalibr/plugin/list"
	"os"
)

func ScanImage(imageName string) error {
	layerCfg := layerimg.DefaultConfig()
	img, err := layerimg.FromLocalDockerImage(imageName, layerCfg)
	if err != nil {
		fmt.Printf("local image loading failed: %s\n", err)
		img, err = layerimg.FromRemoteName(imageName, layerCfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "remote image loading failed: %s\n", err)
			return err
		}
	}

	pluginConf := &config.PluginConfig{}
	pluginNames := []string{"default"}
	// capabilities := &plugin.Capabilities{
	// 	OS: plugin.OSAny,
	// }
	plugins, err := pluglist.FromNames([]string{"all"}, pluginConf)
	// plugins, err := pl.FromCapabilities(capabilities, pluginConf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "plugins configuration failed %s : %s\n", pluginNames, err)
		return err
	}

	for _, plugin := range plugins {
		fmt.Printf("plugin %s configured, reqs: %v\n", plugin.Name(), plugin.Requirements())
	}

	fmt.Printf("%d plugins configured\n", len(plugins))

	scanCfg := &scalibr.ScanConfig{
		Plugins: plugins,
	}

	scanRes, err := scalibr.New().ScanContainer(context.Background(), img, scanCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scan failed: %s\n", err)
		return err
	}

	for _, pkg := range scanRes.Inventory.Packages {
		fmt.Printf(
			"package: {path: %s, name: %s, version: %s}, PURL type: %s, plugins: %s\n",
			pkg.Location.Descriptor.File.Path,
			pkg.Name,
			pkg.Version,
			pkg.PURLType,
			pkg.Plugins,
		)
	}

	return nil
}
