# Demo serverless sam golang with local hot reloading

## Notes

**Deployment**

```bash
aws lambda create-function \
--profile work \
--region ap-southeast-2 \
--function-name main \
--memory 128 \
--runtime go1.x \
--zip-file fileb:///Users/ben/go/src/github.com/benrowe/hello/main.zip \
--handler main
```