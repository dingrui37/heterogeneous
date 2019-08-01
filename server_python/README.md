1. Compile calculte.proto 
```python -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. calculate.proto ```
2. Write Server code
3. Build docker image
docker build -t 