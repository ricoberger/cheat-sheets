# Cheat Sheets

My personal cheat sheets and cheat sheets generator, which generates cheat sheets from JSON files. Inspired by [vim.rtorr.com](https://vim.rtorr.com).

```sh
go build -o ./generator .

for file in $(ls ./data); do
  ./generator data/$file > out/${file//json/html}
done
```

### Formatting

- **Code:** To add a code section a command can be prefixed using `||| `. The code will be rendered after the commands description.
- **Image:** To add an image a command can be prefixed using `!!! ` followed by the image url, e.g. `!!! assets/flux-gitops-toolkit.png`
