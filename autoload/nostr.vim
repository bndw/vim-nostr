function! nostr#Post()
    let content = join(getline(1, '$'), '\n')
    let resp = system("printf  '" . content . "' | vim-nostr -post")
    echo resp
endfunction
