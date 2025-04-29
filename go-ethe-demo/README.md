- Sign up for a free account at Infura ( https://infura.io )
- Create a new project
- Copy your project ID
- Replace "YOUR-PROJECT-ID" in the code above with your actual project ID

```bash
curl --url https://sepolia.infura.io/v3/40a17052b04b43ad884503a6fcf9e3bc \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'
```
we can see the result like this:
```json
{"jsonrpc":"2.0","id":1,"result":"0x7d6dc1"}% 
```