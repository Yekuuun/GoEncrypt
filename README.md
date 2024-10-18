```

                    ____       _____                             _   
                   / ___| ___ | ____|_ __   ___ _ __ _   _ _ __ | |_ 
                  | |  _ / _ \|  _| | '_ \ / __| '__| | | | '_ \| __|
                  | |_| | (_) | |___| | | | (__| |  | |_| | |_) | |_ 
                   \____|\___/|_____|_| |_|\___|_|   \__, | .__/ \__|
                                                     |___/|_|
                  ---simple cli file encryption tool written in GO---

```

This repository contains a base file encryption tool basically for learning GO development. I also choose this project to provide a simple and free way to protect your personnal data.

## Implementation

GoEncrypt use RSA & AES cyphering methods for maximal security using process as follow : 

- Generating RSA keys (pub & private)
- Generating AES for each file encrypted
- Using public key to cipher AES key & adding them into ciphered file
- Using private key to uncypher content of the file


## Assets

<img src="https://github.com/Yekuuun/GoEncrypt/blob/main/assets/cyphering-file.png"></img>

---

**NOTE :**

**This project is very simple & contains base features for learning GO lang.**
