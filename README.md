# AzureCTX

Quickly switch between Azure subscriptions, similar to
[kubectx](https://github.com/ahmetb/kubectx)

## Requirements

[fzf](https://github.com/junegunn/fzf). Just make sure it's installed and in
your `$PATH`

[az](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli). Make sure
it's also installed and in your `$PATH`

## Installation

## Usage

Launch fzf subscription picker with:

```bash
$ azurectx pick
...
# Or just run with no options
$ azurectx
...
```

List all subscriptions with:

```bash
$ azurectx list
Test Subscription 1
Test Subscription 2
```

List current subscription with:

```bash
$ azurectx current
Test Subscription 2
```

Set a specific subscription non-interactively with:

```bash
$ azurectx set "Test Subscription 1"
Switched to 'TestSubscription 1'
```
