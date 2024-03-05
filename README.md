# kubemerge
![kubemerge logo](logo/logo.png)

`kubemerge` is a command-line tool that simplifies the process of merging a standalone kubeconfig file into the default kubeconfig located at `$HOME/.kube/config`. This tool is especially useful for Kubernetes administrators and users who manage multiple clusters and need to consolidate their configurations seamlessly.

## Features

- Easy merging of standalone kubeconfig files into the default kubeconfig.
- Minimal dependencies, ensuring a lightweight and fast operation.
- User-friendly command-line interface.

## Installation

To install `kubemerge`, follow these steps:

1. Clone the repository:
```
git clone https://github.com/stianfro/kubemerge.git
```

2. Build the project (ensure you have Go installed):
```
cd kubemerge
go build -o kubemerge
```

## Usage

To merge a kubeconfig file, simply run:

```
./kubemerge /path/to/your/standalone/kubeconfig
```

Replace `/path/to/your/standalone/kubeconfig` with the actual path to your kubeconfig file.

## Contributing

Contributions to `kubemerge` are welcome! Feel free to open issues or submit pull requests.

## License

`kubemerge` is released under the MIT License. See the [LICENSE](LICENSE) file for more details.

