cd ../ontology
./ontology contract deploy --needstore --code ../ontology-tool/params/shardasset/xshardasstdemo.avm.str --name demo --version 1 --author test --email test@test.com --desc 'xshard asset test' --gasprice 0 --gaslimit 20000000
./ontology contract deploy --needstore --code ../ontology-tool/params/shardcontract/xshardcaller.avm.str --name demo --version 1 --author test --email test@test.com --desc 'xshard asset test' --gasprice 0 --gaslimit 20000000
./ontology contract deploy --needstore --code ../ontology-tool/params/shardcontract/xshardcallee.avm.str --name demo --version 1 --author test --email test@test.com --desc 'xshard asset test' --gasprice 0 --gaslimit 20000000
sleep 10
# xshard asset init
./ontology contract invoke --address a4b7710d5286352c4a23a51eef5e87d02c590617 --params string:init,[int:0] --gasprice 0 --gaslimit 3000000
# xshard callee init
./ontology contract invoke --address 88493a7ebae5e0431854f3f0b7e8f791f5e2d089 --params string:init,[int:0] --gasprice 0 --gaslimit 3000000
sleep 10
# caller dependent callee
./ontology contract invoke --address 0d9b18e994330002d823cd9809543cada9a5a2c1 --params string:init,[int:0] --gasprice 0 --gaslimit 3000000

cd ../ontology-tool
./main -t TransferOntMultiSign
./main -t TransferFromOngMultiSign
./main -t ShardInit
./main -t ShardCreate
./main -t ShardConfig
./main -t ShardPeerApply
./main -t ShardPeerApprove
./main -t ShardPeerJoin
./main -t ShardActivate
./main -t ShardAssetInit