### box

box is a tool for encrypting and decrypting files using secretbox

### usage

Generate a secretbox key
```
    $ box generate
    J5s5xEl3TskRnwHUZI3lYPMf7RFRwvD2PVoHWxMF0m0=
```

Set BOX_KEY environment variable
```
    $ export BOX_KEY='J5s5xEl3TskRnwHUZI3lYPMf7RFRwvD2PVoHWxMF0m0='
```
[NB: `J5s5xEl3TskRnwHUZI3lYPMf7RFRwvD2PVoHWxMF0m0=` is a demonstration key, do not use this key]

Seal a file
```
    $ box seal plain.txt > cipher.txt
```

Open a sealed file
```
    $ box open cipher.txt > plain.txt
```

### install

```
    $ go get github.com/voutasaurus/box
```

### contributions

PRs welcome
