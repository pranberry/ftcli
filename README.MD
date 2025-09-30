# FTCLI
## A file transfer utility
- Used to transfer a file over a network.
- Using symetric keys and ChaCha20-Poly1305 for encryption. KDF is Argon2.
- Built for memory efficiency
    - Files are loaded into memory in 4096-bytes chunks. This helps reduce the memory load from large files
    - MACs are calculated for each individual chunk

## Command syntax:
- `ftcli -h`
- Send file:     `ftcli -rip=192.168.60.70 -role=send -file=/loc/to/file -pass=passw0rd`
- Recieve file:  `ftcli -role=recv -pass=passw0rd`

## Technical criteria
- Files are encrypted with password during transfer
- The reciever has the option to decline incoming transfer
- Able to send a file, and validate file on recieving end (hashes must match)
- The file is never fully loaded into memeory:
    - While no loading in memory needed if using a stream cipher like with AES or ChaCha20 (alone, no MAC)
    - However, when adding a Poly1305 MAC, since the ciphertext must be computed before sending, memory loading is required
    - Effective pattern is to send the file in chunks, computing the MAC at chunk level
- Only use the Go stdlib or extended stdlib
    - Extended lib must be used while implementing ChaCha20-Poly1305

## Next steps
- if the download is interupted, you should be able to resume from there you left off.
- for large files, offer compression
- improve the syntax of the command:
    - `ftcli -send -rip 192.168.60.70 -f /loc/to/file` (prompted for pass)
- parallel processing for super-large files? files over X chunks in size
