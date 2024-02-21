# vim-nostr

Experimental tooling for publishing nostr events from vim.

Draft your post in vim and when you're ready to publish, run `:NostrPost`.
This will create and publish a kind 1 note to your relays.

## Setup

1. Create the required config file at `~/.config/nostr/config.json`:

> Note: The config file is the same format used by [algia](https://github.com/mattn/algia?tab=readme-ov-file#configuration)
and [noscl](https://github.com/fiatjaf/noscl).


```json
{
  "relays": {
    "wss://relay.damus.io": {
      "read": true,
      "write": true
    }
  },
  "privatekey": "<hex priv key>"
}
```

2. Install the Go program that bridges vim and nostr:

```
go install github.com/bndw/vim-nostr@latest
```

3. Install the vim-nostr vim plugin. We're going to use `vim-plug` but feel free to use other plugin managers instead. 

```
curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
git clone https://github.com/bndw/vim-nostr.git ~/.vim/plugged/vim-nostr
```

Create a `~/.vimrc` with the following:

```
call plug#begin()
Plug 'bndw/vim-nostr'
call plug#end()
```
