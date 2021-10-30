# AzureCTX

Quickly switch between Azure subscriptions, similar to
[kubectx](https://github.com/ahmetb/kubectx)

## Requirements

The only external requirement currently is
[fzf](https://github.com/junegunn/fzf). Just make sure it's installed and in
your `$PATH`

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
...
```

List current subscription with:

```bash
$ azurectx -c
...
```

Set a specific subscription non-interactively with:

```bash
$ azurectx "Subscription Name"
...
```
