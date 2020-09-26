# Fuzz

![](https://img.shields.io/badge/Language-Go-brightgreen)

Fuzz fabric (v1.1 and v2.0) smart contracts  based on Beego.

Before doing this, please make sure these are installed：

```
Hyperledger Fabric v1.1
Hyperledger Fabric v2.0
Beego
Go-fuzz
Go v1.12.1
```

Notes：

```
fabric v1.1 Path：$GOPATH/src/github.com/hyperledger/fabric
fabric v2.0 Path：$GOPATH/src/github.com/hyperledger/fabric2.0
aab Folder Path：$GOPATH/src/github.com/hyperledger/fabric2.0/integration/chaincode/
```

If there are any bugs, I am pleased to accept your advice.

Step 1：

```
cd $GOPATH
mkdir example
cd example
git clone https://github.com/doporg/Fuzz.git
```

Step 2：

```
cd server
bee run
```

Step 3：

Front-end

```
git clone https://github.com/doporg/dop-web.git
```

Step 4：

```
npm install
npm start
```

Step 5：

```
http://localhost:3000/#/fuzz
```

### Smart contracts to be tested

The smart contracts to be tested must meet the following conditions：

```
1.Fabric v1.1 or v2.0
2.package is named as “test”
3.Chaincode object in v1.1 smart contracts is named as "Chaincode", and in v2.0 is named as "SimpleChaincode"
```

### Form format

**Test Case**

The initial fuzzing corpus is given in json format, the key is the function name of the smart contract to be tested, and the value is the parameter of the function function.

For example, a smart contract to be tested contains two function functions named `addIngInfo` and `getIngInfo`. The addIngInfo function contains three parameters named FoodID, IngID, and IngName, and the getIngInfo function contains one parameter named FoodID. Then the format of the test case is as follows:

```
{"addIngInfo":{"FoodID":"001","IngID":"aa","IngName":"bb"},"getFoodInfo":{"FoodID":"001"},"getIngInfo":{"FoodID":"001"}}
```

**Version**

Enter 1.1 to fuzz the v1.1 fabric  smart contracts, and 2.0 to fuzz the v2.0 fabric  smart contracts.

**Smart Contract**

Upload the smart contract to be tested.

**Button**

```
start：start fuzzing
stop：stop fuzzing，and get the result data (recommend to execute the test for more than 10 minutes)
Download：download test results, including crashers、corpus、suppressions and output.log
```

### Dockerfile

**Packed image**

```
docker pull alliac0901/fuzz-image:1.2
docker pull alliac0901/fuzz-image:1.0 (Only support v2.0 fabric smart contracts)
docker pull registry.dop.clsaa.com/dop/fuzz-image:1.2
```

**Pack local image**

Dependent packages（put in the root directory）

| Folder name |                     Description                      |
| :---------: | :--------------------------------------------------: |
|   dvyukov   | go-fuzz tool：https://github.com/dvyukov/go-fuzz.git |
|  fabric2.0  |               Hyperledger Fabric v2.0                |
|     go      |                      go v1.12.1                      |
|     bin     |     Executable file：bee、go-fuzz、go-fuzz-build     |

Step 1：

```
cd server
docker build -t fuzz-image .
```

Step 2：

```
docker images
//Check if there is a mirror named fuzz-image
```

Step 3：

```
docker run -it --name test-instance -p 8080:8080 fuzz-image
```

End
