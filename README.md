# dotfair

Dotfair aims to be a simple, easy to use, and fast way to measure your terraform deployments environmental impact.

> Note: This is a work in progress, and is not yet ready for production use. Please feel free to contribute, or open issues. 

## Features
At the moment dotfair only provide an estimation for `aws_instance` resources. The following features are planned for the future:

- [ ] Support for more resources (aws_ebs_volume, aws_elasticache_cluster, etc).
- [ ] Cost estimation of resources
- [ ] Region factor to account for different power grids and energy sources.
- [ ] Support for more cloud providers (Azure, GCP, etc)
- [ ] Support for more metrics (CO2, etc)
- [ ] Support for more output formats (CSV, etc)
- [ ] Support for other exporters (Prometheus, SQL, etc)
- [ ] Support for other input formats (HCL, etc)
- [ ] Support for other input sources (Terraform Cloud, etc)

If you have any other ideas, please open an issue.

## Installation
There are a few ways to install dotfair, the easiest is to use the docker image, but you can also build from source.
### Using Docker
```bash
docker run -it --rm -v $(pwd)/<your_terraform_folder>:/app/terraform ghcr.io/dotfair-opensource/dotfair:latest  ./dotfair run
```

### From source
```bash
git clone https://github.com/dotfair-opensource/dotfair.git
cd dotfair
make build
./bin/dotfair run -w <your_terraform_folder> -f <output_format>
```

## Usage

### Root Command
```bash
dotfair is a CLI to measure your cloud environmental footprint before you deploy

Usage:
  dotfair [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  run         Run a scan of your terraform code

Flags:
      --config string   Specify a config file, there are no defaults
  -h, --help            help for dotfair

Use "dotfair [command] --help" for more information about a command.
```

### Run
```bash
Usage:
  dotfair run [flags]

Flags:
  -f, --format string      Specify the format of the output, either human-readable or json (default "human-readable")
  -h, --help               help for run
  -v, --verbose            Enable verbose output
  -w, --workspace string   Specify the folder to scan containing terraform code (default "./terraform")

Global Flags:
      --config string   Specify a config file, there are no defaults
```

## Contributing
If you would like to contribute to this project, please feel free to open a PR or an issue.

## License
This project is licensed under the terms of the MIT license. See the [LICENSE](LICENSE) file.