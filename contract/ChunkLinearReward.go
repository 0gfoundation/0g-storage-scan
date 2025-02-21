// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ChunkLinearRewardMetaData contains all meta data concerning the ChunkLinearReward contract.
var ChunkLinearRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"releaseSeconds_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"minerId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributeReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PARAMS_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"minerID\",\"type\":\"bytes32\"}],\"name\":\"claimMineReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beforeLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chargedSectors\",\"type\":\"uint256\"}],\"name\":\"fillReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstRewardableChunk\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mine_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"market\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mine\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"}],\"name\":\"rewardDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"lockedReward\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimableReward\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"distributedReward\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"startTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdate\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeRateBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseReward_\",\"type\":\"uint256\"}],\"name\":\"setBaseReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"setServiceFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdrawPayments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a03461006f57601f6120bf38819003918201601f19168301916001600160401b038311848410176100745780849260209460405283398101031261006f5751608052604051612034908161008b8239608051818181610431015281816110170152818161127501526116860152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146110a1575080630373a23a1461107e5780630a539a191461105f578063158ef93e1461103a5780632129593114610fff578063248a9ca314610fc95780632f2ff15d14610f8857806331b3eb9414610eff57806336568abe14610ea1578063485cc95514610c9b57806359e96700146109b957806361d027b31461099257806376ad03bc146109745780637f1b5e431461095657806380f556051461092f5780639010d07c146108ea57806391d148541461089d57806399f4b251146108765780639b1d309114610853578063a217fddf14610837578063b15d20da146107fc578063b3b30c1a14610687578063b7a3c04c14610391578063c057511114610373578063ca15c87314610349578063d547741f146102ff578063e2982c2114610243578063ed88c68e14610225578063f0f44260146101d95763f301af421461016b57600080fd5b346101d65760203660031901126101d657604060a091600435815260056020522064ffffffffff60018254920154604051926001600160801b038116845260801c60208401526001600160801b0381166040840152818160801c16606084015260a81c166080820152f35b80fd5b50346101d65760203660031901126101d6576001600160a01b036101fb611193565b610203611318565b1673ffffffffffffffffffffffffffffffffffffffff19600954161760095580f35b50806003193601126101d65761023d346006546111fd565b60065580f35b50346101d65760203660031901126101d65761025d611193565b60206001600160a01b036024818554169360405194859384927fe3a9db1a0000000000000000000000000000000000000000000000000000000084521660048301525afa9081156102f45782916102ba575b602082604051908152f35b90506020813d6020116102ec575b816102d5602093836111db565b810103126102e8576020915051386102af565b5080fd5b3d91506102c8565b6040513d84823e3d90fd5b50346101d65760403660031901126101d65761034560043561031f611178565b9061034061033b82600052600160205260016040600020015490565b61138a565b61140b565b5080f35b50346101d65760203660031901126101d65760406020916004358152600283522054604051908152f35b50346101d657806003193601126101d6576020600854604051908152f35b50346101d65760603660031901126101d6576004356103ae611178565b906004546001600160a01b031633146103c6906112a3565b808352600560205260408320916040516103df816111a9565b8354936001600160801b0385168252602082019460801c85526001015460408201906001600160801b0381168252606083018160801c64ffffffffff168152608084019160a81c64ffffffffff1682527f000000000000000000000000000000000000000000000000000000000000000061045a908561198e565b6001600160801b03168085516001600160801b03169061047991611655565b6001600160801b0316855287516001600160801b03169061049991611441565b6001600160801b03811688524264ffffffffff16835260011c6f7fffffffffffffffffffffffffffffff16968781516001600160801b0316906104db91611655565b6001600160801b031681528784516001600160801b0316906104fc91611441565b6001600160801b039081168552878a52600560205260408a2086519251608090811b6fffffffffffffffffffffffffffffffff1916938316939093178155945160019095018054935194519490921b64ffffffffff60801b1694167fffffffffffff0000000000000000000000000000000000000000000000000000909216919091179290921760a89190911b64ffffffffff60a81b1617905561059f90611675565b600654936105c0918086111561067b576105ba9080926111fd565b9461130b565b600655826105cc578380f35b6001600160a01b031690836001600160a01b03815416803b156102e8578185916024604051809481937ff340fa010000000000000000000000000000000000000000000000000000000083528960048401525af180156102f457610662575b50506040519283527ff099de0c527342bff1397d08daae31c3de66104f48886094fbe9126a21ebf4b0602060443594a43880808380f35b8161066c916111db565b61067757833861062b565b8380fd5b506105ba8580926111fd565b50346101d657806003193601126101d657610400815b6106b067ffffffffffffffff831661120a565b80151590816107f2575b50156106fc57508060011b6801fffffffffffffffe67fffffffffffffffe8216911681036106e8579061069d565b602483634e487b7160e01b81526011600452fd5b67ffffffffffffffff821667ffffffffffffffff821690808210156107dd5781900367ffffffffffffffff81116107c95760011c677fffffffffffffff16019067ffffffffffffffff82116107b55767ffffffffffffffff8216916107608361120a565b80151590816107ab575b501561079857505060010167ffffffffffffffff81116106e85767ffffffffffffffff90915b9190506106fc565b90925067ffffffffffffffff9150610790565b905042113861076a565b602484634e487b7160e01b81526011600452fd5b602485634e487b7160e01b81526011600452fd5b60208367ffffffffffffffff60405191168152f35b90504211386106ba565b50346101d657806003193601126101d65760206040517fb9d69e0ca90be54a40811e436234a7f7908b87ff2bec27e64f878b166da8e8e58152f35b50346101d657806003193601126101d657602090604051908152f35b50346101d65760203660031901126101d65761086d611318565b60043560085580f35b50346101d657806003193601126101d65760206001600160a01b0360045416604051908152f35b50346101d65760403660031901126101d6576001600160a01b0360406108c1611178565b9260043581526001602052209116600052602052602060ff604060002054166040519015158152f35b50346101d65760403660031901126101d6576001600160a01b0361091f60209260043581526002845260406024359120611960565b90549060031b1c16604051908152f35b50346101d657806003193601126101d65760206001600160a01b0360035416604051908152f35b50346101d657806003193601126101d6576020600654604051908152f35b50346101d657806003193601126101d6576020600754604051908152f35b50346101d657806003193601126101d65760206001600160a01b0360095416604051908152f35b5060403660031901126101d657600435602435906109e36001600160a01b036003541633146112a3565b6127106109f2600854346112ee565b0480610bd7575b610a03903461130b565b630200000092838202828104851483151715610b7157610a2682610a2d92611301565b91846111fd565b908415610b3d57848406850393858511610b715784610a4b916111fd565b60008615610b9b5750859004600019810192818411610bc3576000198101818111610baf578715610b3d578790069160018301809311610baf57610a8f838361130b565b60008915610b9b5750889004916001830190818411610b8757898202918083048b1490151715610b71571496858303610adf575050505050610adc935084526005602052604084206115b7565b80f35b91939550919388526005602052610afa6040892091866112ee565b908715610b3d5787610b0d920490611461565b818110610b53575090610b2b918652600560205260408620926112ee565b928015610b3d57610adc9304906115b7565b634e487b7160e01b600052601260045260246000fd5b8060019188526005602052610b6b8560408a20611461565b01610b0d565b634e487b7160e01b600052601160045260246000fd5b60248b634e487b7160e01b81526011600452fd5b80634e487b7160e01b602492526012600452fd5b602489634e487b7160e01b81526011600452fd5b602488634e487b7160e01b81526011600452fd5b6001600160a01b0360095416814710610c6f578480808481945af13d15610c6a573d67ffffffffffffffff8111610c565760405190610c20601f8201601f1916602001836111db565b81528560203d92013e5b6109f9576004847f1425ea42000000000000000000000000000000000000000000000000000000008152fd5b602486634e487b7160e01b81526041600452fd5b610c2a565b6024857fcd78605900000000000000000000000000000000000000000000000000000000815230600452fd5b50346101d65760403660031901126101d657610cb5611193565b610cbd611178565b825460ff8160a01c16610e38576001600160a01b0392740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff859316178555610d19336116bd565b610e1d575b610d2733611769565b610de2575b1673ffffffffffffffffffffffffffffffffffffffff1960035416176003551673ffffffffffffffffffffffffffffffffffffffff1960045416176004556040516104988082019082821067ffffffffffffffff831117610dce57908291611b678339039082f08015610dc1576001600160a01b031673ffffffffffffffffffffffffffffffffffffffff1982541617815580f35b50604051903d90823e3d90fd5b602484634e487b7160e01b81526041600452fd5b7fb9d69e0ca90be54a40811e436234a7f7908b87ff2bec27e64f878b166da8e8e585526002602052610e173360408720611a1d565b50610d2c565b8480526002602052610e323360408720611a1d565b50610d1e565b608460405162461bcd60e51b8152602060048201526024808201527f5a67496e697469616c697a61626c653a20616c726561647920696e697469616c60448201527f697a6564000000000000000000000000000000000000000000000000000000006064820152fd5b50346101d65760403660031901126101d657610ebb611178565b336001600160a01b03821603610ed7576103459060043561140b565b6004827f6697b232000000000000000000000000000000000000000000000000000000008152fd5b50346101d65760203660031901126101d657610f19611193565b816001600160a01b0381541691823b156102e8576001600160a01b03602483928360405196879485937f51cff8d90000000000000000000000000000000000000000000000000000000085521660048401525af18015610dc157610f7a5780f35b610f83916111db565b388180f35b50346101d65760403660031901126101d657610345600435610fa8611178565b90610fc461033b82600052600160205260016040600020015490565b6113d1565b50346101d65760203660031901126101d6576020610ff7600435600052600160205260016040600020015490565b604051908152f35b50346101d657806003193601126101d65760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b50346101d657806003193601126101d65760ff6020915460a01c166040519015158152f35b50346101d65760203660031901126101d6576020610ff760043561120a565b50346101d65760203660031901126101d657611098611318565b60043560075580f35b9050346102e85760203660031901126102e8576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361117457602092507f5a05180f000000000000000000000000000000000000000000000000000000008114908115611117575b5015158152f35b7f7965db0b0000000000000000000000000000000000000000000000000000000081149150811561114a575b5038611110565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438611143565b8280fd5b602435906001600160a01b038216820361118e57565b600080fd5b600435906001600160a01b038216820361118e57565b60a0810190811067ffffffffffffffff8211176111c557604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff8211176111c557604052565b91908201809211610b7157565b60005260056020526040600020600160405191611226836111a9565b80546001600160801b038116845260801c60208401520154906001600160801b0382166040820152608064ffffffffff8084831c169384606085015260a81c16910152801561129d5761129a907f0000000000000000000000000000000000000000000000000000000000000000906111fd565b90565b50600090565b156112aa57565b606460405162461bcd60e51b815260206004820152601f60248201527f53656e64657220646f6573206e6f742068617665207065726d697373696f6e006044820152fd5b81810292918115918404141715610b7157565b8115610b3d570490565b91908203918211610b7157565b3360009081527f74e899ea3b8c951f087f4e3bf156263af04db888284b8793254674b266200203602052604090205460ff161561135157565b63e2517d3f60e01b600052336004527fb9d69e0ca90be54a40811e436234a7f7908b87ff2bec27e64f878b166da8e8e560245260446000fd5b80600052600160205260406000206001600160a01b03331660005260205260ff60406000205416156113b95750565b63e2517d3f60e01b6000523360045260245260446000fd5b6113db8282611837565b91826113e657505090565b6114079160005260026020526001600160a01b036040600020911690611a1d565b5090565b61141582826118d1565b918261142057505090565b6114079160005260026020526001600160a01b036040600020911690611a93565b906001600160801b03809116911601906001600160801b038211610b7157565b6001600160801b03821161157357600181019064ffffffffff825460801c1661152f576001600160801b0361149e8161152d951682845416611441565b166fffffffffffffffffffffffffffffffff1982541617905564ffffffffff4216906114f482829074ffffffffff000000000000000000000000000000001964ffffffffff60801b83549260801b169116179055565b907fffffffffffff0000000000ffffffffffffffffffffffffffffffffffffffffff64ffffffffff60a81b83549260a81b169116179055565b565b606460405162461bcd60e51b815260206004820152602060248201527f526577617264206974656d20686173206265656e20696e697469616c697a65646044820152fd5b606460405162461bcd60e51b815260206004820152600f60248201527f526577617264206f766572666c6f7700000000000000000000000000000000006044820152fd5b9190916001600160801b03831161157357600181019264ffffffffff845460801c1661152f576115f46001600160801b0380921682845416611441565b166fffffffffffffffffffffffffffffffff198254161790556116145750565b61152d9064ffffffffff4216906114f482829074ffffffffff000000000000000000000000000000001964ffffffffff60801b83549260801b169116179055565b906001600160801b03809116911603906001600160801b038211610b7157565b64ffffffffff60606116ab920151167f0000000000000000000000000000000000000000000000000000000000000000906111fd565b4210156116b85760075490565b600090565b6001600160a01b03811660009081527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49602052604090205460ff1661129d576001600160a01b031660008181527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4960205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b6001600160a01b03811660009081527f74e899ea3b8c951f087f4e3bf156263af04db888284b8793254674b266200203602052604090205460ff1661129d576001600160a01b031660008181527f74e899ea3b8c951f087f4e3bf156263af04db888284b8793254674b26620020360205260408120805460ff191660011790553391907fb9d69e0ca90be54a40811e436234a7f7908b87ff2bec27e64f878b166da8e8e5907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b80600052600160205260406000206001600160a01b03831660005260205260ff60406000205416156000146118ca5780600052600160205260406000206001600160a01b0383166000526020526040600020600160ff198254161790556001600160a01b03339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b80600052600160205260406000206001600160a01b03831660005260205260ff604060002054166000146118ca5780600052600160205260406000206001600160a01b038316600052602052604060002060ff1981541690556001600160a01b03339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4600190565b80548210156119785760005260206000200190600090565b634e487b7160e01b600052603260045260246000fd5b9064ffffffffff608083015116156118ca576119fb6001600160801b036119c2816020860151168260408701511690611441565b16916119f66119f064ffffffffff60606119e6876001600160801b038a51166111fd565b970151164261130b565b856112ee565b611301565b91808311611a15575b508082106118ca5761129a9161130b565b915038611a04565b6001810190826000528160205260406000205415600014611a8b578054680100000000000000008110156111c557611a76611a5f826001879401855584611960565b819391549060031b91821b91600019901b19161790565b90555491600052602052604060002055600190565b505050600090565b9060018201918160005282602052604060002054801515600014611b5d576000198101818111610b71578254600019810191908211610b7157818103611b26575b50505080548015611b10576000190190611aee8282611960565b8154906000199060031b1b191690555560005260205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b611b46611b36611a5f9386611960565b90549060031b1c92839286611960565b905560005283602052604060002055388080611ad4565b5050505060009056fe608080604052346075573315605f5760008054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a361041d908161007b8239f35b631e4fbdf760e01b600052600060045260246000fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806351cff8d914610265578063715018a6146101ff5780638da5cb5b146101d8578063e3a9db1a1461019b578063f2fde38b146100e95763f340fa011461005e57600080fd5b60203660031901126100e4576004356001600160a01b0381168091036100e4576100866103a5565b80600052600160205260406000208054903482018092116100ce57557f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c46020604051348152a2005b634e487b7160e01b600052601160045260246000fd5b600080fd5b346100e45760203660031901126100e4576004356001600160a01b0381168091036100e4576101166103a5565b801561016c576001600160a01b036000548273ffffffffffffffffffffffffffffffffffffffff19821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f1e4fbdf700000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b346100e45760203660031901126100e4576004356001600160a01b0381168091036100e45760005260016020526020604060002054604051908152f35b346100e45760003660031901126100e45760206001600160a01b0360005416604051908152f35b346100e45760003660031901126100e4576102186103a5565b60006001600160a01b03815473ffffffffffffffffffffffffffffffffffffffff1981168355167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346100e45760203660031901126100e4576004356001600160a01b0381168091036100e4576102926103a5565b6000818152600160205260408120805491905547811161037757600080808084865af13d15610372573d67ffffffffffffffff811161035c5760405190601f8101601f19908116603f0116820167ffffffffffffffff81118382101761035c576040528152600060203d92013e5b156103325760207f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d591604051908152a2005b7f1425ea420000000000000000000000000000000000000000000000000000000060005260046000fd5b634e487b7160e01b600052604160045260246000fd5b610300565b7fcd786059000000000000000000000000000000000000000000000000000000006000523060045260246000fd5b6001600160a01b036000541633036103b957565b7f118cdaa7000000000000000000000000000000000000000000000000000000006000523360045260246000fdfea2646970667358221220a0ffa33dd5fd2929ed957f32f39dc88cb96bc633038e78cab0796821a63db3e764736f6c634300081b0033a264697066735822122098b1ffd33933c564d2b17ffd98290c49e4adb17152334c589b529f3b5c13c8f064736f6c634300081b0033",
}

// ChunkLinearRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use ChunkLinearRewardMetaData.ABI instead.
var ChunkLinearRewardABI = ChunkLinearRewardMetaData.ABI

// ChunkLinearRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChunkLinearRewardMetaData.Bin instead.
var ChunkLinearRewardBin = ChunkLinearRewardMetaData.Bin

// DeployChunkLinearReward deploys a new Ethereum contract, binding an instance of ChunkLinearReward to it.
func DeployChunkLinearReward(auth *bind.TransactOpts, backend bind.ContractBackend, releaseSeconds_ *big.Int) (common.Address, *types.Transaction, *ChunkLinearReward, error) {
	parsed, err := ChunkLinearRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChunkLinearRewardBin), backend, releaseSeconds_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChunkLinearReward{ChunkLinearRewardCaller: ChunkLinearRewardCaller{contract: contract}, ChunkLinearRewardTransactor: ChunkLinearRewardTransactor{contract: contract}, ChunkLinearRewardFilterer: ChunkLinearRewardFilterer{contract: contract}}, nil
}

// ChunkLinearReward is an auto generated Go binding around an Ethereum contract.
type ChunkLinearReward struct {
	ChunkLinearRewardCaller     // Read-only binding to the contract
	ChunkLinearRewardTransactor // Write-only binding to the contract
	ChunkLinearRewardFilterer   // Log filterer for contract events
}

// ChunkLinearRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChunkLinearRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChunkLinearRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChunkLinearRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChunkLinearRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChunkLinearRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChunkLinearRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChunkLinearRewardSession struct {
	Contract     *ChunkLinearReward // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ChunkLinearRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChunkLinearRewardCallerSession struct {
	Contract *ChunkLinearRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ChunkLinearRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChunkLinearRewardTransactorSession struct {
	Contract     *ChunkLinearRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ChunkLinearRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChunkLinearRewardRaw struct {
	Contract *ChunkLinearReward // Generic contract binding to access the raw methods on
}

// ChunkLinearRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChunkLinearRewardCallerRaw struct {
	Contract *ChunkLinearRewardCaller // Generic read-only contract binding to access the raw methods on
}

// ChunkLinearRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChunkLinearRewardTransactorRaw struct {
	Contract *ChunkLinearRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChunkLinearReward creates a new instance of ChunkLinearReward, bound to a specific deployed contract.
func NewChunkLinearReward(address common.Address, backend bind.ContractBackend) (*ChunkLinearReward, error) {
	contract, err := bindChunkLinearReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearReward{ChunkLinearRewardCaller: ChunkLinearRewardCaller{contract: contract}, ChunkLinearRewardTransactor: ChunkLinearRewardTransactor{contract: contract}, ChunkLinearRewardFilterer: ChunkLinearRewardFilterer{contract: contract}}, nil
}

// NewChunkLinearRewardCaller creates a new read-only instance of ChunkLinearReward, bound to a specific deployed contract.
func NewChunkLinearRewardCaller(address common.Address, caller bind.ContractCaller) (*ChunkLinearRewardCaller, error) {
	contract, err := bindChunkLinearReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardCaller{contract: contract}, nil
}

// NewChunkLinearRewardTransactor creates a new write-only instance of ChunkLinearReward, bound to a specific deployed contract.
func NewChunkLinearRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*ChunkLinearRewardTransactor, error) {
	contract, err := bindChunkLinearReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardTransactor{contract: contract}, nil
}

// NewChunkLinearRewardFilterer creates a new log filterer instance of ChunkLinearReward, bound to a specific deployed contract.
func NewChunkLinearRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*ChunkLinearRewardFilterer, error) {
	contract, err := bindChunkLinearReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardFilterer{contract: contract}, nil
}

// bindChunkLinearReward binds a generic wrapper to an already deployed contract.
func bindChunkLinearReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChunkLinearRewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChunkLinearReward *ChunkLinearRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChunkLinearReward.Contract.ChunkLinearRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChunkLinearReward *ChunkLinearRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.ChunkLinearRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChunkLinearReward *ChunkLinearRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.ChunkLinearRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChunkLinearReward *ChunkLinearRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChunkLinearReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChunkLinearReward *ChunkLinearRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChunkLinearReward *ChunkLinearRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ChunkLinearReward.Contract.DEFAULTADMINROLE(&_ChunkLinearReward.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ChunkLinearReward.Contract.DEFAULTADMINROLE(&_ChunkLinearReward.CallOpts)
}

// PARAMSADMINROLE is a free data retrieval call binding the contract method 0xb15d20da.
//
// Solidity: function PARAMS_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCaller) PARAMSADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "PARAMS_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PARAMSADMINROLE is a free data retrieval call binding the contract method 0xb15d20da.
//
// Solidity: function PARAMS_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardSession) PARAMSADMINROLE() ([32]byte, error) {
	return _ChunkLinearReward.Contract.PARAMSADMINROLE(&_ChunkLinearReward.CallOpts)
}

// PARAMSADMINROLE is a free data retrieval call binding the contract method 0xb15d20da.
//
// Solidity: function PARAMS_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) PARAMSADMINROLE() ([32]byte, error) {
	return _ChunkLinearReward.Contract.PARAMSADMINROLE(&_ChunkLinearReward.CallOpts)
}

// BaseReward is a free data retrieval call binding the contract method 0x76ad03bc.
//
// Solidity: function baseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) BaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "baseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseReward is a free data retrieval call binding the contract method 0x76ad03bc.
//
// Solidity: function baseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) BaseReward() (*big.Int, error) {
	return _ChunkLinearReward.Contract.BaseReward(&_ChunkLinearReward.CallOpts)
}

// BaseReward is a free data retrieval call binding the contract method 0x76ad03bc.
//
// Solidity: function baseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) BaseReward() (*big.Int, error) {
	return _ChunkLinearReward.Contract.BaseReward(&_ChunkLinearReward.CallOpts)
}

// FirstRewardableChunk is a free data retrieval call binding the contract method 0xb3b30c1a.
//
// Solidity: function firstRewardableChunk() view returns(uint64)
func (_ChunkLinearReward *ChunkLinearRewardCaller) FirstRewardableChunk(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "firstRewardableChunk")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstRewardableChunk is a free data retrieval call binding the contract method 0xb3b30c1a.
//
// Solidity: function firstRewardableChunk() view returns(uint64)
func (_ChunkLinearReward *ChunkLinearRewardSession) FirstRewardableChunk() (uint64, error) {
	return _ChunkLinearReward.Contract.FirstRewardableChunk(&_ChunkLinearReward.CallOpts)
}

// FirstRewardableChunk is a free data retrieval call binding the contract method 0xb3b30c1a.
//
// Solidity: function firstRewardableChunk() view returns(uint64)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) FirstRewardableChunk() (uint64, error) {
	return _ChunkLinearReward.Contract.FirstRewardableChunk(&_ChunkLinearReward.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ChunkLinearReward.Contract.GetRoleAdmin(&_ChunkLinearReward.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ChunkLinearReward.Contract.GetRoleAdmin(&_ChunkLinearReward.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ChunkLinearReward.Contract.GetRoleMember(&_ChunkLinearReward.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ChunkLinearReward.Contract.GetRoleMember(&_ChunkLinearReward.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ChunkLinearReward.Contract.GetRoleMemberCount(&_ChunkLinearReward.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ChunkLinearReward.Contract.GetRoleMemberCount(&_ChunkLinearReward.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ChunkLinearReward.Contract.HasRole(&_ChunkLinearReward.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ChunkLinearReward.Contract.HasRole(&_ChunkLinearReward.CallOpts, role, account)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardSession) Initialized() (bool, error) {
	return _ChunkLinearReward.Contract.Initialized(&_ChunkLinearReward.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Initialized() (bool, error) {
	return _ChunkLinearReward.Contract.Initialized(&_ChunkLinearReward.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Market(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "market")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardSession) Market() (common.Address, error) {
	return _ChunkLinearReward.Contract.Market(&_ChunkLinearReward.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Market() (common.Address, error) {
	return _ChunkLinearReward.Contract.Market(&_ChunkLinearReward.CallOpts)
}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Mine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "mine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardSession) Mine() (common.Address, error) {
	return _ChunkLinearReward.Contract.Mine(&_ChunkLinearReward.CallOpts)
}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Mine() (common.Address, error) {
	return _ChunkLinearReward.Contract.Mine(&_ChunkLinearReward.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Payments(opts *bind.CallOpts, dest common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "payments", dest)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) Payments(dest common.Address) (*big.Int, error) {
	return _ChunkLinearReward.Contract.Payments(&_ChunkLinearReward.CallOpts, dest)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Payments(dest common.Address) (*big.Int, error) {
	return _ChunkLinearReward.Contract.Payments(&_ChunkLinearReward.CallOpts, dest)
}

// ReleaseSeconds is a free data retrieval call binding the contract method 0x21295931.
//
// Solidity: function releaseSeconds() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) ReleaseSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "releaseSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseSeconds is a free data retrieval call binding the contract method 0x21295931.
//
// Solidity: function releaseSeconds() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) ReleaseSeconds() (*big.Int, error) {
	return _ChunkLinearReward.Contract.ReleaseSeconds(&_ChunkLinearReward.CallOpts)
}

// ReleaseSeconds is a free data retrieval call binding the contract method 0x21295931.
//
// Solidity: function releaseSeconds() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) ReleaseSeconds() (*big.Int, error) {
	return _ChunkLinearReward.Contract.ReleaseSeconds(&_ChunkLinearReward.CallOpts)
}

// RewardDeadline is a free data retrieval call binding the contract method 0x0a539a19.
//
// Solidity: function rewardDeadline(uint256 pricingIndex) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) RewardDeadline(opts *bind.CallOpts, pricingIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "rewardDeadline", pricingIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardDeadline is a free data retrieval call binding the contract method 0x0a539a19.
//
// Solidity: function rewardDeadline(uint256 pricingIndex) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) RewardDeadline(pricingIndex *big.Int) (*big.Int, error) {
	return _ChunkLinearReward.Contract.RewardDeadline(&_ChunkLinearReward.CallOpts, pricingIndex)
}

// RewardDeadline is a free data retrieval call binding the contract method 0x0a539a19.
//
// Solidity: function rewardDeadline(uint256 pricingIndex) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) RewardDeadline(pricingIndex *big.Int) (*big.Int, error) {
	return _ChunkLinearReward.Contract.RewardDeadline(&_ChunkLinearReward.CallOpts, pricingIndex)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(uint128 lockedReward, uint128 claimableReward, uint128 distributedReward, uint40 startTime, uint40 lastUpdate)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Rewards(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LockedReward      *big.Int
	ClaimableReward   *big.Int
	DistributedReward *big.Int
	StartTime         *big.Int
	LastUpdate        *big.Int
}, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "rewards", arg0)

	outstruct := new(struct {
		LockedReward      *big.Int
		ClaimableReward   *big.Int
		DistributedReward *big.Int
		StartTime         *big.Int
		LastUpdate        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockedReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ClaimableReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DistributedReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastUpdate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(uint128 lockedReward, uint128 claimableReward, uint128 distributedReward, uint40 startTime, uint40 lastUpdate)
func (_ChunkLinearReward *ChunkLinearRewardSession) Rewards(arg0 *big.Int) (struct {
	LockedReward      *big.Int
	ClaimableReward   *big.Int
	DistributedReward *big.Int
	StartTime         *big.Int
	LastUpdate        *big.Int
}, error) {
	return _ChunkLinearReward.Contract.Rewards(&_ChunkLinearReward.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(uint128 lockedReward, uint128 claimableReward, uint128 distributedReward, uint40 startTime, uint40 lastUpdate)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Rewards(arg0 *big.Int) (struct {
	LockedReward      *big.Int
	ClaimableReward   *big.Int
	DistributedReward *big.Int
	StartTime         *big.Int
	LastUpdate        *big.Int
}, error) {
	return _ChunkLinearReward.Contract.Rewards(&_ChunkLinearReward.CallOpts, arg0)
}

// ServiceFeeRateBps is a free data retrieval call binding the contract method 0xc0575111.
//
// Solidity: function serviceFeeRateBps() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) ServiceFeeRateBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "serviceFeeRateBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFeeRateBps is a free data retrieval call binding the contract method 0xc0575111.
//
// Solidity: function serviceFeeRateBps() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) ServiceFeeRateBps() (*big.Int, error) {
	return _ChunkLinearReward.Contract.ServiceFeeRateBps(&_ChunkLinearReward.CallOpts)
}

// ServiceFeeRateBps is a free data retrieval call binding the contract method 0xc0575111.
//
// Solidity: function serviceFeeRateBps() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) ServiceFeeRateBps() (*big.Int, error) {
	return _ChunkLinearReward.Contract.ServiceFeeRateBps(&_ChunkLinearReward.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ChunkLinearReward.Contract.SupportsInterface(&_ChunkLinearReward.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ChunkLinearReward.Contract.SupportsInterface(&_ChunkLinearReward.CallOpts, interfaceId)
}

// TotalBaseReward is a free data retrieval call binding the contract method 0x7f1b5e43.
//
// Solidity: function totalBaseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) TotalBaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "totalBaseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseReward is a free data retrieval call binding the contract method 0x7f1b5e43.
//
// Solidity: function totalBaseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) TotalBaseReward() (*big.Int, error) {
	return _ChunkLinearReward.Contract.TotalBaseReward(&_ChunkLinearReward.CallOpts)
}

// TotalBaseReward is a free data retrieval call binding the contract method 0x7f1b5e43.
//
// Solidity: function totalBaseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) TotalBaseReward() (*big.Int, error) {
	return _ChunkLinearReward.Contract.TotalBaseReward(&_ChunkLinearReward.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardSession) Treasury() (common.Address, error) {
	return _ChunkLinearReward.Contract.Treasury(&_ChunkLinearReward.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Treasury() (common.Address, error) {
	return _ChunkLinearReward.Contract.Treasury(&_ChunkLinearReward.CallOpts)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 minerID) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) ClaimMineReward(opts *bind.TransactOpts, pricingIndex *big.Int, beneficiary common.Address, minerID [32]byte) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "claimMineReward", pricingIndex, beneficiary, minerID)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 minerID) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) ClaimMineReward(pricingIndex *big.Int, beneficiary common.Address, minerID [32]byte) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.ClaimMineReward(&_ChunkLinearReward.TransactOpts, pricingIndex, beneficiary, minerID)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 minerID) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) ClaimMineReward(pricingIndex *big.Int, beneficiary common.Address, minerID [32]byte) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.ClaimMineReward(&_ChunkLinearReward.TransactOpts, pricingIndex, beneficiary, minerID)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) Donate() (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.Donate(&_ChunkLinearReward.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) Donate() (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.Donate(&_ChunkLinearReward.TransactOpts)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 chargedSectors) payable returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) FillReward(opts *bind.TransactOpts, beforeLength *big.Int, chargedSectors *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "fillReward", beforeLength, chargedSectors)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 chargedSectors) payable returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) FillReward(beforeLength *big.Int, chargedSectors *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.FillReward(&_ChunkLinearReward.TransactOpts, beforeLength, chargedSectors)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 chargedSectors) payable returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) FillReward(beforeLength *big.Int, chargedSectors *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.FillReward(&_ChunkLinearReward.TransactOpts, beforeLength, chargedSectors)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.GrantRole(&_ChunkLinearReward.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.GrantRole(&_ChunkLinearReward.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) Initialize(opts *bind.TransactOpts, market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "initialize", market_, mine_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) Initialize(market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.Initialize(&_ChunkLinearReward.TransactOpts, market_, mine_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) Initialize(market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.Initialize(&_ChunkLinearReward.TransactOpts, market_, mine_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.RenounceRole(&_ChunkLinearReward.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.RenounceRole(&_ChunkLinearReward.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.RevokeRole(&_ChunkLinearReward.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.RevokeRole(&_ChunkLinearReward.TransactOpts, role, account)
}

// SetBaseReward is a paid mutator transaction binding the contract method 0x0373a23a.
//
// Solidity: function setBaseReward(uint256 baseReward_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) SetBaseReward(opts *bind.TransactOpts, baseReward_ *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "setBaseReward", baseReward_)
}

// SetBaseReward is a paid mutator transaction binding the contract method 0x0373a23a.
//
// Solidity: function setBaseReward(uint256 baseReward_) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) SetBaseReward(baseReward_ *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetBaseReward(&_ChunkLinearReward.TransactOpts, baseReward_)
}

// SetBaseReward is a paid mutator transaction binding the contract method 0x0373a23a.
//
// Solidity: function setBaseReward(uint256 baseReward_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) SetBaseReward(baseReward_ *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetBaseReward(&_ChunkLinearReward.TransactOpts, baseReward_)
}

// SetServiceFeeRate is a paid mutator transaction binding the contract method 0x9b1d3091.
//
// Solidity: function setServiceFeeRate(uint256 bps) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) SetServiceFeeRate(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "setServiceFeeRate", bps)
}

// SetServiceFeeRate is a paid mutator transaction binding the contract method 0x9b1d3091.
//
// Solidity: function setServiceFeeRate(uint256 bps) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) SetServiceFeeRate(bps *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetServiceFeeRate(&_ChunkLinearReward.TransactOpts, bps)
}

// SetServiceFeeRate is a paid mutator transaction binding the contract method 0x9b1d3091.
//
// Solidity: function setServiceFeeRate(uint256 bps) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) SetServiceFeeRate(bps *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetServiceFeeRate(&_ChunkLinearReward.TransactOpts, bps)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) SetTreasury(opts *bind.TransactOpts, treasury_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "setTreasury", treasury_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) SetTreasury(treasury_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetTreasury(&_ChunkLinearReward.TransactOpts, treasury_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) SetTreasury(treasury_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetTreasury(&_ChunkLinearReward.TransactOpts, treasury_)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) WithdrawPayments(opts *bind.TransactOpts, payee common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "withdrawPayments", payee)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) WithdrawPayments(payee common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.WithdrawPayments(&_ChunkLinearReward.TransactOpts, payee)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) WithdrawPayments(payee common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.WithdrawPayments(&_ChunkLinearReward.TransactOpts, payee)
}

// ChunkLinearRewardDistributeRewardIterator is returned from FilterDistributeReward and is used to iterate over the raw logs and unpacked data for DistributeReward events raised by the ChunkLinearReward contract.
type ChunkLinearRewardDistributeRewardIterator struct {
	Event *ChunkLinearRewardDistributeReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChunkLinearRewardDistributeRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChunkLinearRewardDistributeReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChunkLinearRewardDistributeReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChunkLinearRewardDistributeRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChunkLinearRewardDistributeRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChunkLinearRewardDistributeReward represents a DistributeReward event raised by the ChunkLinearReward contract.
type ChunkLinearRewardDistributeReward struct {
	PricingIndex *big.Int
	Beneficiary  common.Address
	MinerId      [32]byte
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDistributeReward is a free log retrieval operation binding the contract event 0xf099de0c527342bff1397d08daae31c3de66104f48886094fbe9126a21ebf4b0.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, bytes32 indexed minerId, uint256 amount)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) FilterDistributeReward(opts *bind.FilterOpts, pricingIndex []*big.Int, beneficiary []common.Address, minerId [][32]byte) (*ChunkLinearRewardDistributeRewardIterator, error) {

	var pricingIndexRule []interface{}
	for _, pricingIndexItem := range pricingIndex {
		pricingIndexRule = append(pricingIndexRule, pricingIndexItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var minerIdRule []interface{}
	for _, minerIdItem := range minerId {
		minerIdRule = append(minerIdRule, minerIdItem)
	}

	logs, sub, err := _ChunkLinearReward.contract.FilterLogs(opts, "DistributeReward", pricingIndexRule, beneficiaryRule, minerIdRule)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardDistributeRewardIterator{contract: _ChunkLinearReward.contract, event: "DistributeReward", logs: logs, sub: sub}, nil
}

// WatchDistributeReward is a free log subscription operation binding the contract event 0xf099de0c527342bff1397d08daae31c3de66104f48886094fbe9126a21ebf4b0.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, bytes32 indexed minerId, uint256 amount)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) WatchDistributeReward(opts *bind.WatchOpts, sink chan<- *ChunkLinearRewardDistributeReward, pricingIndex []*big.Int, beneficiary []common.Address, minerId [][32]byte) (event.Subscription, error) {

	var pricingIndexRule []interface{}
	for _, pricingIndexItem := range pricingIndex {
		pricingIndexRule = append(pricingIndexRule, pricingIndexItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var minerIdRule []interface{}
	for _, minerIdItem := range minerId {
		minerIdRule = append(minerIdRule, minerIdItem)
	}

	logs, sub, err := _ChunkLinearReward.contract.WatchLogs(opts, "DistributeReward", pricingIndexRule, beneficiaryRule, minerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChunkLinearRewardDistributeReward)
				if err := _ChunkLinearReward.contract.UnpackLog(event, "DistributeReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributeReward is a log parse operation binding the contract event 0xf099de0c527342bff1397d08daae31c3de66104f48886094fbe9126a21ebf4b0.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, bytes32 indexed minerId, uint256 amount)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) ParseDistributeReward(log types.Log) (*ChunkLinearRewardDistributeReward, error) {
	event := new(ChunkLinearRewardDistributeReward)
	if err := _ChunkLinearReward.contract.UnpackLog(event, "DistributeReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChunkLinearRewardRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleAdminChangedIterator struct {
	Event *ChunkLinearRewardRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChunkLinearRewardRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChunkLinearRewardRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChunkLinearRewardRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChunkLinearRewardRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChunkLinearRewardRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChunkLinearRewardRoleAdminChanged represents a RoleAdminChanged event raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ChunkLinearRewardRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ChunkLinearReward.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardRoleAdminChangedIterator{contract: _ChunkLinearReward.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ChunkLinearRewardRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ChunkLinearReward.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChunkLinearRewardRoleAdminChanged)
				if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) ParseRoleAdminChanged(log types.Log) (*ChunkLinearRewardRoleAdminChanged, error) {
	event := new(ChunkLinearRewardRoleAdminChanged)
	if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChunkLinearRewardRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleGrantedIterator struct {
	Event *ChunkLinearRewardRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChunkLinearRewardRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChunkLinearRewardRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChunkLinearRewardRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChunkLinearRewardRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChunkLinearRewardRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChunkLinearRewardRoleGranted represents a RoleGranted event raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ChunkLinearRewardRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChunkLinearReward.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardRoleGrantedIterator{contract: _ChunkLinearReward.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ChunkLinearRewardRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChunkLinearReward.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChunkLinearRewardRoleGranted)
				if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) ParseRoleGranted(log types.Log) (*ChunkLinearRewardRoleGranted, error) {
	event := new(ChunkLinearRewardRoleGranted)
	if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChunkLinearRewardRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleRevokedIterator struct {
	Event *ChunkLinearRewardRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChunkLinearRewardRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChunkLinearRewardRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChunkLinearRewardRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChunkLinearRewardRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChunkLinearRewardRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChunkLinearRewardRoleRevoked represents a RoleRevoked event raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ChunkLinearRewardRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChunkLinearReward.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardRoleRevokedIterator{contract: _ChunkLinearReward.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ChunkLinearRewardRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChunkLinearReward.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChunkLinearRewardRoleRevoked)
				if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) ParseRoleRevoked(log types.Log) (*ChunkLinearRewardRoleRevoked, error) {
	event := new(ChunkLinearRewardRoleRevoked)
	if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
