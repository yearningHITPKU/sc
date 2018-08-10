# sc
1.将yjs.jar放到你的代码执行目录下，命令行下执行：java -classpath yjs.jar com.yancloud.sc.ContractManager

2.签署合约调用StartContract接口，参数例子："{\\"type\\":\\"Data\\",\\"id\\":\\"656564\\"}"

3.执行合约调用ExecContract接口，参数例子："{\\"arg\\":\\"http://www.baidu.com\",\"contractID\":\"656564\"}"