// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PreimageOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageLengths\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageParts\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1002,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimagePartOk\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\"},{\"astId\":1003,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"zeroHashes\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_bytes32)16_storage\"},{\"astId\":1004,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposals\",\"offset\":0,\"slot\":\"19\",\"type\":\"t_array(t_struct(LargePreimageProposalKeys)1009_storage)dyn_storage\"},{\"astId\":1005,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalBranches\",\"offset\":0,\"slot\":\"20\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_bytes32)16_storage))\"},{\"astId\":1006,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalMetadata\",\"offset\":0,\"slot\":\"21\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010))\"},{\"astId\":1007,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalParts\",\"offset\":0,\"slot\":\"22\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1008,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalBlocks\",\"offset\":0,\"slot\":\"23\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_uint64)dyn_storage))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_bytes32)16_storage\":{\"encoding\":\"inplace\",\"label\":\"bytes32[16]\",\"numberOfBytes\":\"512\",\"base\":\"t_bytes32\"},\"t_array(t_struct(LargePreimageProposalKeys)1009_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct PreimageOracle.LargePreimageProposalKeys[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(LargePreimageProposalKeys)1009_storage\"},\"t_array(t_uint64)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint64[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint64\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_bytes32)16_storage))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bytes32[16]))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_array(t_bytes32)16_storage)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_uint64)dyn_storage))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e uint64[]))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_array(t_uint64)dyn_storage)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e LPPMetaData))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_array(t_bytes32)16_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32[16])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_bytes32)16_storage\"},\"t_mapping(t_uint256,t_array(t_uint64)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint64[])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_uint64)dyn_storage\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_mapping(t_uint256,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bytes32\"},\"t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e LPPMetaData)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_userDefinedValueType(LPPMetaData)1010\"},\"t_struct(LargePreimageProposalKeys)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"struct PreimageOracle.LargePreimageProposalKeys\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_userDefinedValueType(LPPMetaData)1010\":{\"encoding\":\"inplace\",\"label\":\"LPPMetaData\",\"numberOfBytes\":\"32\"}}}"

var PreimageOracleStorageLayout = new(solc.StorageLayout)

var PreimageOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106101ae5760003560e01c80639d53a648116100ee578063dd24f9bf11610097578063ec5efcbc11610071578063ec5efcbc14610458578063f3f480d91461046b578063faf37bc714610491578063fef2b4ed146104a457600080fd5b8063dd24f9bf146103f7578063e03110e11461041d578063e15926111461044557600080fd5b8063b4801e61116100c8578063b4801e61146103c9578063d18534b5146103dc578063da35c664146103ef57600080fd5b80639d53a648146103495780639d7e87691461038b578063b2e67ba81461039e57600080fd5b806361238bde1161015b5780637ac54767116101355780637ac54767146102b95780638542cf50146102cc578063882856ef1461030a5780638dc4be111461033657600080fd5b806361238bde146102505780636551927b1461027b5780637917de1d146102a657600080fd5b80633909af5c1161018c5780633909af5c146102205780634d52b4c91461023557806352f0f3ad1461023d57600080fd5b8063013cf08b146101b35780630359a563146101f75780632055b36b14610218575b600080fd5b6101c66101c1366004612a0c565b6104c4565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152015b60405180910390f35b61020a610205366004612a4e565b610509565b6040519081526020016101ee565b61020a601081565b61023361022e366004612c49565b610641565b005b61020a61088f565b61020a61024b366004612d35565b6108aa565b61020a61025e366004612d70565b600160209081526000928352604080842090915290825290205481565b61020a610289366004612a4e565b601560209081526000928352604080842090915290825290205481565b6102336102b4366004612dd4565b61097f565b61020a6102c7366004612a0c565b610e6f565b6102fa6102da366004612d70565b600260209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016101ee565b61031d610318366004612e70565b610e86565b60405167ffffffffffffffff90911681526020016101ee565b610233610344366004612ea3565b610ee0565b61020a610357366004612a4e565b73ffffffffffffffffffffffffffffffffffffffff9091166000908152601760209081526040808320938352929052205490565b610233610399366004612eef565b610fdb565b61020a6103ac366004612a4e565b601660209081526000928352604080842090915290825290205481565b61020a6103d7366004612e70565b611223565b6102336103ea366004612c49565b611255565b60135461020a565b7f000000000000000000000000000000000000000000000000000000000000000061020a565b61043061042b366004612d70565b61160c565b604080519283526020830191909152016101ee565b610233610453366004612ea3565b6116fd565b610233610466366004612f7b565b611805565b7f000000000000000000000000000000000000000000000000000000000000000061020a565b61023361049f366004613014565b61197f565b61020a6104b2366004612a0c565b60006020819052908152604090205481565b601381815481106104d457600080fd5b60009182526020909120600290910201805460019091015473ffffffffffffffffffffffffffffffffffffffff909116915082565b73ffffffffffffffffffffffffffffffffffffffff82166000908152601560209081526040808320848452909152812054819061054c9060601c63ffffffff1690565b63ffffffff16905060005b601081101561063957816001166001036105df5773ffffffffffffffffffffffffffffffffffffffff85166000908152601460209081526040808320878452909152902081601081106105ac576105ac613050565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250610620565b82600382601081106105f3576105f3613050565b01546040805160208101939093528201526060016040516020818303038152906040528051906020012092505b60019190911c9080610631816130ae565b915050610557565b505092915050565b600061064d8a8a610509565b905061067086868360208b013561066b6106668d6130e6565b611b92565b611bd2565b801561068e575061068e838383602088013561066b6106668a6130e6565b6106c4576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8660400135886040516020016106da91906131b5565b6040516020818303038152906040528051906020012014610727576040517f1968a90200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83602001358760200135600161073d91906131f3565b14610774576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107bc88610782868061320b565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611c3392505050565b6107c588611d8e565b8360400135886040516020016107db91906131b5565b6040516020818303038152906040528051906020012003610828576040517f9843145b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505073ffffffffffffffffffffffffffffffffffffffff9590951660009081526015602090815260408083209683529590529390932080547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117905550505050565b600161089d60106002613392565b6108a7919061339e565b81565b60006108b6868661262a565b90506108c38360086131f3565b8211806108d05750602083115b15610907576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602081815260c085901b82526008959095528251828252600286526040808320858452875280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558484528752808320948352938652838220558181529384905292205592915050565b606081156109985761099186866126d7565b90506109d2565b85858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293505050505b3360009081526014602090815260408083208b845290915280822081516102008101928390529160109082845b8154815260200190600101908083116109ff57505050505090506000601560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b81526020019081526020016000205490506000610a808260601c63ffffffff1690565b63ffffffff169050333214610ac1576040517fba092d1600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ad18260801c63ffffffff1690565b63ffffffff16600003610b10576040517f87138d5c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b1a8260c01c90565b67ffffffffffffffff1615610b5b576040517f475a253500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b898114610b94576040517f60f95d5a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ba189898d8886612750565b83516020850160888204881415608883061715610bc6576307b1daf16000526004601cfd5b60405160c8810160405260005b83811015610c76578083018051835260208101516020840152604081015160408401526060810151606084015260808101516080840152508460888301526088810460051b8b013560a883015260c882206001860195508560005b610200811015610c6b576001821615610c4b5782818b0152610c6b565b8981015160009081526020938452604090209260019290921c9101610c2e565b505050608801610bd3565b50505050600160106002610c8a9190613392565b610c94919061339e565b811115610ccd576040517f6229572300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d42610ce08360401c63ffffffff1690565b610cf09063ffffffff168a6131f3565b60401b7fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff606084901b167fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff8516171790565b91508415610dcf5777ffffffffffffffffffffffffffffffffffffffffffffffff82164260c01b179150610d7c8260801c63ffffffff1690565b63ffffffff16610d928360401c63ffffffff1690565b63ffffffff1614610dcf576040517f7b1dafd100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526014602090815260408083208e84529091529020610df590846010612982565b50503360008181526017602090815260408083208e8452825280832080546001810182559084528284206004820401805460039092166008026101000a67ffffffffffffffff818102199093164390931602919091179055928252601581528282209c82529b909b52909920989098555050505050505050565b60038160108110610e7f57600080fd5b0154905081565b60176020528260005260406000206020528160005260406000208181548110610eae57600080fd5b906000526020600020906004918282040191900660080292509250509054906101000a900467ffffffffffffffff1681565b604435600080600883018610610efe5763fe2549876000526004601cfd5b60c083901b60805260888386823786600882030151915060206000858360025afa905080610f2b57600080fd5b50600080517effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0400000000000000000000000000000000000000000000000000000000000000178082526002602090815260408084208a8552825280842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558385528252808420998452988152888320939093558152908190529490942055505050565b7f0000000000000000000000000000000000000000000000000000000000000000421015611035576040517f299f254900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080603087600037602060006030600060025afa8061105d5763f91129696000526004601cfd5b6000517effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f010000000000000000000000000000000000000000000000000000000000000017608081815260a08c905260c08b905260308a60e037603088609083013760008060c083600a5afa9250826110df576309bde3396000526004601cfd5b602886106110f55763fe2549876000526004601cfd5b6000602882015278200000000000000000000000000000000000000000000000008152600881018b905285810151935060308a8237603081018c9052605090207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0500000000000000000000000000000000000000000000000000000000000000176000818152600260209081526040808320898452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155848452825280832089845282528083208790558383528282529182902055519094507f7fdb715eccae6f912623299cb64da860cdd085cad5d1f65b5151ce89e892302a9250611210915084815260200190565b60405180910390a1505050505050505050565b6014602052826000526040600020602052816000526040600020816010811061124b57600080fd5b0154925083915050565b73ffffffffffffffffffffffffffffffffffffffff891660009081526015602090815260408083208b845290915290205467ffffffffffffffff8116156112c8576040517fc334f06900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006112f38260c01c90565b6113079067ffffffffffffffff164261339e565b1161133e576040517f55d4cbf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061134a8b8b610509565b905061136387878360208c013561066b6106668e6130e6565b80156113815750611381848483602089013561066b6106668b6130e6565b6113b7576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8760400135896040516020016113cd91906131b5565b604051602081830303815290604052805190602001201461141a576040517f1968a90200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84602001358860200135600161143091906131f3565b1415806114625750600161144a8360601c63ffffffff1690565b61145491906133b5565b63ffffffff16856020013514155b15611499576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114a789610782878061320b565b6114b089611d8e565b60006114bb8a6128a3565b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f020000000000000000000000000000000000000000000000000000000000000017905060006115128460a01c63ffffffff1690565b67ffffffffffffffff169050600160026000848152602001908152602001600020600083815260200190815260200160002060006101000a81548160ff021916908315150217905550601660008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008d815260200190815260200160002054600160008481526020019081526020016000206000838152602001908152602001600020819055506115e48460801c63ffffffff1690565b60009283526020839052604090922063ffffffff909216909155505050505050505050505050565b6000828152600260209081526040808320848452909152812054819060ff16611695576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015260640160405180910390fd5b50600083815260208181526040909120546116b18160086131f3565b6116bc8560206131f3565b106116da57836116cd8260086131f3565b6116d7919061339e565b91505b506000938452600160209081526040808620948652939052919092205492909150565b60443560008060088301861061171b5763fe2549876000526004601cfd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b60006118118686610509565b905061182a838383602088013561066b6106668a6130e6565b611860576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208401351561189c576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6118a46129c0565b6118b281610782878061320b565b6118bb81611d8e565b8460400135816040516020016118d191906131b5565b604051602081830303815290604052805190602001200361191e576040517f9843145b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505073ffffffffffffffffffffffffffffffffffffffff9290921660009081526015602090815260408083209383529290522080547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117905550565b3332146119b8576040517fba092d1600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119c38160086133da565b63ffffffff168263ffffffff1610611a07576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000008163ffffffff161015611a67576040517f7b1dafd100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000818152601560209081526040808320878452825280832080547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1660a09790971b7fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff169690961760809590951b9490941790945582518084019093529082529181019283526013805460018101825592525160029091027f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a0908101805473ffffffffffffffffffffffffffffffffffffffff9093167fffffffffffffffffffffffff00000000000000000000000000000000000000009093169290921790915590517f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a09190910155565b6000816000015182602001518360400151604051602001611bb593929190613402565b604051602081830303815290604052805190602001209050919050565b60008160005b6010811015611c26578060051b880135600186831c1660018114611c0b5760008481526020839052604090209350611c1c565b600082815260208590526040902093505b5050600101611bd8565b5090931495945050505050565b6088815114611c4157600080fd5b6020810160208301611cc2565b8260031b8201518060001a8160011a60081b178160021a60101b8260031a60181b17178160041a60201b8260051a60281b178260061a60301b8360071a60381b1717179050611cbc81611ca7868560059190911b015190565b1867ffffffffffffffff16600586901b840152565b50505050565b611cce60008383611c4e565b611cda60018383611c4e565b611ce660028383611c4e565b611cf260038383611c4e565b611cfe60048383611c4e565b611d0a60058383611c4e565b611d1660068383611c4e565b611d2260078383611c4e565b611d2e60088383611c4e565b611d3a60098383611c4e565b611d46600a8383611c4e565b611d52600b8383611c4e565b611d5e600c8383611c4e565b611d6a600d8383611c4e565b611d76600e8383611c4e565b611d82600f8383611c4e565b611cbc60108383611c4e565b6040805178010000000000008082800000000000808a8000000080008000602082015279808b00000000800000018000000080008081800000000000800991810191909152788a00000000000000880000000080008009000000008000000a60608201527b8000808b800000000000008b8000000000008089800000000000800360808201527f80000000000080028000000000000080000000000000800a800000008000000a60a08201527f800000008000808180000000000080800000000080000001800000008000800860c082015260009060e0016040516020818303038152906040529050602082016020820161250a565b6102808101516101e082015161014083015160a0840151845118189118186102a082015161020083015161016084015160c0850151602086015118189118186102c083015161022084015161018085015160e0860151604087015118189118186102e08401516102408501516101a0860151610100870151606088015118189118186103008501516102608601516101c0870151610120880151608089015118189118188084603f1c611f418660011b67ffffffffffffffff1690565b18188584603f1c611f5c8660011b67ffffffffffffffff1690565b18188584603f1c611f778660011b67ffffffffffffffff1690565b181895508483603f1c611f948560011b67ffffffffffffffff1690565b181894508387603f1c611fb18960011b67ffffffffffffffff1690565b60208b01518b51861867ffffffffffffffff168c5291189190911897508118600181901b603f9190911c18935060c08801518118601481901c602c9190911b1867ffffffffffffffff1660208901526101208801518718602c81901c60149190911b1867ffffffffffffffff1660c08901526102c08801518618600381901c603d9190911b1867ffffffffffffffff166101208901526101c08801518718601981901c60279190911b1867ffffffffffffffff166102c08901526102808801518218602e81901c60129190911b1867ffffffffffffffff166101c089015260408801518618600281901c603e9190911b1867ffffffffffffffff166102808901526101808801518618601581901c602b9190911b1867ffffffffffffffff1660408901526101a08801518518602781901c60199190911b1867ffffffffffffffff166101808901526102608801518718603881901c60089190911b1867ffffffffffffffff166101a08901526102e08801518518600881901c60389190911b1867ffffffffffffffff166102608901526101e08801518218601781901c60299190911b1867ffffffffffffffff166102e089015260808801518718602581901c601b9190911b1867ffffffffffffffff166101e08901526103008801518718603281901c600e9190911b1867ffffffffffffffff1660808901526102a08801518118603e81901c60029190911b1867ffffffffffffffff166103008901526101008801518518600981901c60379190911b1867ffffffffffffffff166102a08901526102008801518118601381901c602d9190911b1867ffffffffffffffff1661010089015260a08801518218601c81901c60249190911b1867ffffffffffffffff1661020089015260608801518518602481901c601c9190911b1867ffffffffffffffff1660a08901526102408801518518602b81901c60159190911b1867ffffffffffffffff1660608901526102208801518618603181901c600f9190911b1867ffffffffffffffff166102408901526101608801518118603681901c600a9190911b1867ffffffffffffffff166102208901525060e08701518518603a81901c60069190911b1867ffffffffffffffff166101608801526101408701518118603d81901c60039190911b1867ffffffffffffffff1660e0880152505067ffffffffffffffff81166101408601525050505050565b61233181611e84565b805160208201805160408401805160608601805160808801805167ffffffffffffffff871986168a188116808c528619851689188216909952831982169095188516909552841988169091188316909152941990921618811690925260a08301805160c0808601805160e0880180516101008a0180516101208c018051861985168a188d16909a528319821686188c16909652801989169092188a169092528619861618881690529219909216909218841690526101408401805161016086018051610180880180516101a08a0180516101c08c0180518619851689188d169099528319821686188c16909652801988169092188a169092528519851618881690529119909116909118841690526101e08401805161020086018051610220880180516102408a0180516102608c0180518619851689188d169099528319821686188c16909652801988169092188a16909252851985161888169052911990911690911884169052610280840180516102a0860180516102c0880180516102e08a0180516103008c0180518619851689188d169099528319821686188c16909652801988169092188a16909252851985161888169052911990911690911884169052600386901b850151901c9081189091168252611cbc565b61251660008284612328565b61252260018284612328565b61252e60028284612328565b61253a60038284612328565b61254660048284612328565b61255260058284612328565b61255e60068284612328565b61256a60078284612328565b61257660088284612328565b61258260098284612328565b61258e600a8284612328565b61259a600b8284612328565b6125a6600c8284612328565b6125b2600d8284612328565b6125be600e8284612328565b6125ca600f8284612328565b6125d660108284612328565b6125e260118284612328565b6125ee60128284612328565b6125fa60138284612328565b61260660148284612328565b61261260158284612328565b61261e60168284612328565b611cbc60178284612328565b7f01000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316176126d0818360408051600093845233602052918152606090922091527effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001790565b9392505050565b60606040519050816020820181810182868337608883068080156127205760888290038501848101848103803687375060806001820353506001845160001a1784538652612737565b608836843760018353608060878401536088850186525b5050505050601f19603f82510116810160405292915050565b60006127628260a01c63ffffffff1690565b67ffffffffffffffff16905060006127808360801c63ffffffff1690565b63ffffffff169050600061279a8460401c63ffffffff1690565b63ffffffff1690506008831080156127b0575080155b156127e45760c082901b6000908152883560085283513382526016602090815260408084208a855290915290912055612899565b600883101580156128025750806127fc60088561339e565b93508310155b8015612816575061281387826131f3565b83105b15612899576000612827828561339e565b9050876128358260206131f3565b10158015612841575085155b15612878576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526016602090815260408083208a845290915290209089013590555b5050505050505050565b6000612926565b66ff00ff00ff00ff8160081c1667ff00ff00ff00ff006128d48360081b67ffffffffffffffff1690565b1617905065ffff0000ffff8160101c1667ffff0000ffff00006129018360101b67ffffffffffffffff1690565b1617905060008160201c61291f8360201b67ffffffffffffffff1690565b1792915050565b6080820151602083019061293e906128aa565b6128aa565b604082015161294c906128aa565b60401b1761296461293960018460059190911b015190565b825160809190911b90612976906128aa565b60c01b17179392505050565b82601081019282156129b0579160200282015b828111156129b0578251825591602001919060010190612995565b506129bc9291506129d8565b5090565b60405180602001604052806129d36129ed565b905290565b5b808211156129bc57600081556001016129d9565b6040518061032001604052806019906020820280368337509192915050565b600060208284031215612a1e57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114612a4957600080fd5b919050565b60008060408385031215612a6157600080fd5b612a6a83612a25565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610320810167ffffffffffffffff81118282101715612acb57612acb612a78565b60405290565b6040516060810167ffffffffffffffff81118282101715612acb57612acb612a78565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612b3b57612b3b612a78565b604052919050565b6000610320808385031215612b5757600080fd5b604051602080820167ffffffffffffffff8382108183111715612b7c57612b7c612a78565b8160405283955087601f880112612b9257600080fd5b612b9a612aa7565b9487019491508188861115612bae57600080fd5b875b86811015612bd65780358381168114612bc95760008081fd5b8452928401928401612bb0565b50909352509295945050505050565b600060608284031215612bf757600080fd5b50919050565b60008083601f840112612c0f57600080fd5b50813567ffffffffffffffff811115612c2757600080fd5b6020830191508360208260051b8501011115612c4257600080fd5b9250929050565b60008060008060008060008060006103e08a8c031215612c6857600080fd5b612c718a612a25565b985060208a01359750612c878b60408c01612b43565b96506103608a013567ffffffffffffffff80821115612ca557600080fd5b612cb18d838e01612be5565b97506103808c0135915080821115612cc857600080fd5b612cd48d838e01612bfd565b90975095506103a08c0135915080821115612cee57600080fd5b612cfa8d838e01612be5565b94506103c08c0135915080821115612d1157600080fd5b50612d1e8c828d01612bfd565b915080935050809150509295985092959850929598565b600080600080600060a08688031215612d4d57600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b60008060408385031215612d8357600080fd5b50508035926020909101359150565b60008083601f840112612da457600080fd5b50813567ffffffffffffffff811115612dbc57600080fd5b602083019150836020828501011115612c4257600080fd5b600080600080600080600060a0888a031215612def57600080fd5b8735965060208801359550604088013567ffffffffffffffff80821115612e1557600080fd5b612e218b838c01612d92565b909750955060608a0135915080821115612e3a57600080fd5b50612e478a828b01612bfd565b90945092505060808801358015158114612e6057600080fd5b8091505092959891949750929550565b600080600060608486031215612e8557600080fd5b612e8e84612a25565b95602085013595506040909401359392505050565b600080600060408486031215612eb857600080fd5b83359250602084013567ffffffffffffffff811115612ed657600080fd5b612ee286828701612d92565b9497909650939450505050565b600080600080600080600060a0888a031215612f0a57600080fd5b8735965060208801359550604088013567ffffffffffffffff80821115612f3057600080fd5b612f3c8b838c01612d92565b909750955060608a0135915080821115612f5557600080fd5b50612f628a828b01612d92565b989b979a50959894979596608090950135949350505050565b600080600080600060808688031215612f9357600080fd5b612f9c86612a25565b945060208601359350604086013567ffffffffffffffff80821115612fc057600080fd5b612fcc89838a01612be5565b94506060880135915080821115612fe257600080fd5b50612fef88828901612bfd565b969995985093965092949392505050565b803563ffffffff81168114612a4957600080fd5b60008060006060848603121561302957600080fd5b8335925061303960208501613000565b915061304760408501613000565b90509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036130df576130df61307f565b5060010190565b6000606082360312156130f857600080fd5b613100612ad1565b823567ffffffffffffffff8082111561311857600080fd5b9084019036601f83011261312b57600080fd5b813560208282111561313f5761313f612a78565b61316f817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011601612af4565b9250818352368183860101111561318557600080fd5b81818501828501376000918301810191909152908352848101359083015250604092830135928101929092525090565b81516103208201908260005b60198110156131ea57825167ffffffffffffffff168252602092830192909101906001016131c1565b50505092915050565b600082198211156132065761320661307f565b500190565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261324057600080fd5b83018035915067ffffffffffffffff82111561325b57600080fd5b602001915036819003821315612c4257600080fd5b600181815b808511156132c957817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156132af576132af61307f565b808516156132bc57918102915b93841c9390800290613275565b509250929050565b6000826132e05750600161338c565b816132ed5750600061338c565b8160018114613303576002811461330d57613329565b600191505061338c565b60ff84111561331e5761331e61307f565b50506001821b61338c565b5060208310610133831016604e8410600b841016171561334c575081810a61338c565b6133568383613270565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156133885761338861307f565b0290505b92915050565b60006126d083836132d1565b6000828210156133b0576133b061307f565b500390565b600063ffffffff838116908316818110156133d2576133d261307f565b039392505050565b600063ffffffff8083168185168083038211156133f9576133f961307f565b01949350505050565b6000845160005b818110156134235760208188018101518583015201613409565b81811115613432576000828501525b509190910192835250602082015260400191905056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(PreimageOracleStorageLayoutJSON), PreimageOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PreimageOracle"] = PreimageOracleStorageLayout
	deployedBytecodes["PreimageOracle"] = PreimageOracleDeployedBin
	immutableReferences["PreimageOracle"] = true
}
