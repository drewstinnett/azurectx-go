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

Launch fzf subscription picket with:

```bash
$ azurectx
...
```

List all subscriptions with:

```bash
$ azurectx -l
Test Subscription 1
Test Subscription 2
```

List current subscription with:

```bash
$ azurectx -c
Test Subscription 2
```

Set a specific subscription non-interactively with:

```bash
$ azurectx "Test Subscription 1"
Switched to 'TestSubscription 1'
```
