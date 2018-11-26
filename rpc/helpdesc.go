package rpc

import "github.com/copernet/copernicus/util"

func HelpExampleCli(methodname string, args ...string) string {
	help := "> coperctl " + methodname
	for _, arg := range args {
		help += " " + arg
	}
	help += "\n"
	return help
}

func HelpExampleRPC(methodname string, args ...string) string {
	help := "> curl --user myusername --data-binary '{\"jsonrpc\": \"1.0\", " +
		"\"id\":\"curltest\", \"method\": \"" + methodname + "\", \"params\": ["
	for index, arg := range args {
		if index == len(args)-1 {
			help += arg
		} else {
			help += arg + ", "
		}
	}
	help += "] }' -H 'content-type: text/plain;' http://127.0.0.1:18334/\n"
	return help
}

// abc
var (
	getexcessiveblockDesc = "getexcessiveblock\n" +
		"\nReturn the excessive block size." +
		"\nResult\n" +
		"  excessiveBlockSize (integer) block size in bytes\n" +
		"\nExamples:\n" +
		HelpExampleCli("getexcessiveblock") +
		HelpExampleRPC("getexcessiveblock")

	setexcessiveblockDesc = "setexcessiveblock blockSize\n" +
		"\nSet the excessive block size. Excessive blocks will not be used " +
		"in the active chain or relayed. This  discourages the propagation " +
		"of blocks that you consider excessively large." +
		"\nResult\n" +
		"  blockSize (integer) excessive block size in bytes\n" +
		"\nExamples:\n" +
		HelpExampleCli("setexcessiveblock") +
		HelpExampleRPC("setexcessiveblock")
)

//blockchain
var (
	getblockchaininfoDesc = "getblockchaininfo\n" +
		"Returns an object containing various state info regarding " +
		"blockchain processing.\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"chain\": \"xxxx\",        (string) current network name as " +
		"defined in BIP70 (main, test, regtest)\n" +
		"  \"blocks\": xxxxxx,         (numeric) the current number of " +
		"blocks processed in the server\n" +
		"  \"headers\": xxxxxx,        (numeric) the current number of " +
		"headers we have validated\n" +
		"  \"bestblockhash\": \"...\", (string) the hash of the currently " +
		"best block\n" +
		"  \"difficulty\": xxxxxx,     (numeric) the current difficulty\n" +
		"  \"mediantime\": xxxxxx,     (numeric) median time for the " +
		"current best block\n" +
		"  \"verificationprogress\": xxxx, (numeric) estimate of " +
		"verification progress [0..1]\n" +
		"  \"chainwork\": \"xxxx\"     (string) total amount of work in " +
		"active chain, in hexadecimal\n" +
		"  \"pruned\": xx,             (boolean) if the blocks are subject " +
		"to pruning\n" +
		"  \"pruneheight\": xxxxxx,    (numeric) lowest-height complete " +
		"block stored\n" +
		"  \"softforks\": [            (array) status of softforks in " +
		"progress\n" +
		"     {\n" +
		"        \"id\": \"xxxx\",        (string) name of softfork\n" +
		"        \"version\": xx,         (numeric) block version\n" +
		"        \"reject\": {            (object) progress toward " +
		"rejecting pre-softfork blocks\n" +
		"           \"status\": xx,       (boolean) true if threshold " +
		"reached\n" +
		"        },\n" +
		"     }, ...\n" +
		"  ],\n" +
		"  \"bip9_softforks\": {          (object) status of BIP9 " +
		"softforks in progress\n" +
		"     \"xxxx\" : {                (string) name of the softfork\n" +
		"        \"status\": \"xxxx\",    (string) one of \"defined\", " +
		"\"started\", \"locked_in\", \"active\", \"failed\"\n" +
		"        \"bit\": xx,             (numeric) the bit (0-28) in the " +
		"block version field used to signal this softfork (only for " +
		"\"started\" status)\n" +
		"        \"startTime\": xx,       (numeric) the minimum median " +
		"time past of a block at which the bit gains its meaning\n" +
		"        \"timeout\": xx,         (numeric) the median time past " +
		"of a block at which the deployment is considered failed if not " +
		"yet locked in\n" +
		"        \"since\": xx            (numeric) height of the first " +
		"block to which the status applies\n" +
		"     }\n" +
		"  }\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getblockchaininfo") +
		HelpExampleRPC("getblockchaininfo")

	getbestblockhashDesc = "getbestblockhash\n" +
		"\nReturns the hash of the best (tip) block in the " +
		"longest blockchain.\n" +
		"\nResult:\n" +
		"\"hex\"      (string) the block hash hex encoded\n" +
		"\nExamples:\n" +
		HelpExampleCli("getbestblockhash") +
		HelpExampleRPC("getbestblockhash")

	getblockcountDesc = "getblockcount\n" +
		"\nReturns the number of blocks in the longest blockchain.\n" +
		"\nResult:\n" +
		"n    (numeric) The current block count\n" +
		"\nExamples:\n" +
		HelpExampleCli("getblockcount") +
		HelpExampleRPC("getblockcount")

	getblockDesc = "getblock \"blockhash\" ( verbose )\n" +
		"\nIf verbose is false, returns a string that is serialized, " +
		"hex-encoded data for block 'hash'.\n" +
		"If verbose is true, returns an Object with information about " +
		"block <hash>.\n" +
		"\nArguments:\n" +
		"1. \"blockhash\"          (string, required) The block hash\n" +
		"2. verbose                (boolean, optional, default=true) true " +
		"for a json object, false for the hex encoded data\n" +
		"\nResult (for verbose = true):\n" +
		"{\n" +
		"  \"hash\" : \"hash\",     (string) the block hash (same as " +
		"provided)\n" +
		"  \"confirmations\" : n,   (numeric) The number of confirmations, " +
		"or -1 if the block is not on the main chain\n" +
		"  \"size\" : n,            (numeric) The block size\n" +
		"  \"height\" : n,          (numeric) The block height or index\n" +
		"  \"version\" : n,         (numeric) The block version\n" +
		"  \"versionHex\" : \"00000000\", (string) The block version " +
		"formatted in hexadecimal\n" +
		"  \"merkleroot\" : \"xxxx\", (string) The merkle root\n" +
		"  \"tx\" : [               (array of string) The transaction ids\n" +
		"     \"transactionid\"     (string) The transaction id\n" +
		"     ,...\n" +
		"  ],\n" +
		"  \"time\" : ttt,          (numeric) The block time in seconds " +
		"since epoch (Jan 1 1970 GMT)\n" +
		"  \"mediantime\" : ttt,    (numeric) The median block time in " +
		"seconds since epoch (Jan 1 1970 GMT)\n" +
		"  \"nonce\" : n,           (numeric) The nonce\n" +
		"  \"bits\" : \"1d00ffff\", (string) The bits\n" +
		"  \"difficulty\" : x.xxx,  (numeric) The difficulty\n" +
		"  \"chainwork\" : \"xxxx\",  (string) Expected number of hashes " +
		"required to produce the chain up to this block (in hex)\n" +
		"  \"previousblockhash\" : \"hash\",  (string) The hash of the " +
		"previous block\n" +
		"  \"nextblockhash\" : \"hash\"       (string) The hash of the " +
		"next block\n" +
		"}\n" +
		"\nResult (for verbose=false):\n" +
		"\"data\"             (string) A string that is serialized, " +
		"hex-encoded data for block 'hash'.\n" +
		"\nExamples:\n" +
		HelpExampleCli("getblock", "\"00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09\"") +
		HelpExampleRPC("getblock", "\"00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09\"")

	getblockhashDesc = "getblockhash height\n" +
		"\nReturns hash of block in best-block-chain at height provided.\n" +
		"\nArguments:\n" +
		"1. height         (numeric, required) The height index\n" +
		"\nResult:\n" +
		"\"hash\"         (string) The block hash\n" +
		"\nExamples:\n" +
		HelpExampleCli("getblockhash", "1000") +
		HelpExampleRPC("getblockhash", "1000")

	getblockheader = "getblockheader \"hash\" ( verbose )\n" +
		"\nIf verbose is false, returns a string that is serialized, " +
		"hex-encoded data for blockheader 'hash'.\n" +
		"If verbose is true, returns an Object with information about " +
		"blockheader <hash>.\n" +
		"\nArguments:\n" +
		"1. \"hash\"          (string, required) The block hash\n" +
		"2. verbose           (boolean, optional, default=true) true for a " +
		"json object, false for the hex encoded data\n" +
		"\nResult (for verbose = true):\n" +
		"{\n" +
		"  \"hash\" : \"hash\",     (string) the block hash (same as " +
		"provided)\n" +
		"  \"confirmations\" : n,   (numeric) The number of confirmations, " +
		"or -1 if the block is not on the main chain\n" +
		"  \"height\" : n,          (numeric) The block height or index\n" +
		"  \"version\" : n,         (numeric) The block version\n" +
		"  \"versionHex\" : \"00000000\", (string) The block version " +
		"formatted in hexadecimal\n" +
		"  \"merkleroot\" : \"xxxx\", (string) The merkle root\n" +
		"  \"time\" : ttt,          (numeric) The block time in seconds " +
		"since epoch (Jan 1 1970 GMT)\n" +
		"  \"mediantime\" : ttt,    (numeric) The median block time in " +
		"seconds since epoch (Jan 1 1970 GMT)\n" +
		"  \"nonce\" : n,           (numeric) The nonce\n" +
		"  \"bits\" : \"1d00ffff\", (string) The bits\n" +
		"  \"difficulty\" : x.xxx,  (numeric) The difficulty\n" +
		"  \"chainwork\" : \"0000...1f3\"     (string) Expected number of " +
		"hashes required to produce the current chain (in hex)\n" +
		"  \"previousblockhash\" : \"hash\",  (string) The hash of the " +
		"previous block\n" +
		"  \"nextblockhash\" : \"hash\",      (string) The hash of the " +
		"next block\n" +
		"}\n" +
		"\nResult (for verbose=false):\n" +
		"\"data\"             (string) A string that is serialized, " +
		"hex-encoded data for block 'hash'.\n" +
		"\nExamples:\n" +
		HelpExampleCli("getblockheader", "\"00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09\"") +
		HelpExampleRPC("getblockheader", "\"00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09\"")

	getchaintipsDesc = "getchaintips\n" +
		"Return information about all known tips in the block tree," +
		" including the main chain as well as orphaned branches.\n" +
		"\nResult:\n" +
		"[\n" +
		"  {\n" +
		"    \"height\": xxxx,         (numeric) height of the chain tip\n" +
		"    \"hash\": \"xxxx\",         (string) block hash of the tip\n" +
		"    \"branchlen\": 0          (numeric) zero for main chain\n" +
		"    \"status\": \"active\"      (string) \"active\" for the main " +
		"chain\n" +
		"  },\n" +
		"  {\n" +
		"    \"height\": xxxx,\n" +
		"    \"hash\": \"xxxx\",\n" +
		"    \"branchlen\": 1          (numeric) length of branch " +
		"connecting the tip to the main chain\n" +
		"    \"status\": \"xxxx\"        (string) status of the chain " +
		"(active, valid-fork, valid-headers, headers-only, invalid)\n" +
		"  }\n" +
		"]\n" +
		"Possible values for status:\n" +
		"1.  \"invalid\"               This branch contains at least one " +
		"invalid block\n" +
		"2.  \"headers-only\"          Not all blocks for this branch are " +
		"available, but the headers are valid\n" +
		"3.  \"valid-headers\"         All blocks are available for this " +
		"branch, but they were never fully validated\n" +
		"4.  \"valid-fork\"            This branch is not part of the " +
		"active chain, but is fully validated\n" +
		"5.  \"active\"                This is the tip of the active main " +
		"chain, which is certainly valid\n" +
		"\nExamples:\n" +
		HelpExampleCli("getchaintips") +
		HelpExampleRPC("getchaintips")

	getchaintxstatsDesc = "getchaintxstats ( nblocks blockhash )\n" +
		"\nCompute statistics about the total number and rate of " +
		"transactions in the chain.\n" +
		"\nArguments:\n" +
		"1. nblocks      (numeric, optional) Size of the window in number " +
		"of blocks (default: one month).\n" +
		"2. \"blockhash\"  (string, optional) The hash of the block that " +
		"ends the window.\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"time\": xxxxx,                (numeric) The timestamp for the " +
		"final block in the window in UNIX format.\n" +
		"  \"txcount\": xxxxx,             (numeric) The total number of " +
		"transactions in the chain up to that point.\n" +
		"  \"window_block_count\": xxxxx,  (numeric) Size of the window in " +
		"number of blocks.\n" +
		"  \"window_tx_count\": xxxxx,     (numeric) The number of " +
		"transactions in the window. Only returned if " +
		"\"window_block_count\" is > 0.\n" +
		"  \"window_interval\": xxxxx,     (numeric) The elapsed time in " +
		"the window in seconds. Only returned if \"window_block_count\" is " +
		"> 0.\n" +
		"  \"txrate\": x.xx,               (numeric) The average rate of " +
		"transactions per second in the window. Only returned if " +
		"\"window_interval\" is > 0.\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getchaintxstats") +
		HelpExampleRPC("getchaintxstats")

	getdifficultyDesc = "getdifficulty\n" +
		"\nReturns the proof-of-work difficulty as a " +
		"multiple of the minimum difficulty.\n" +
		"\nResult:\n" +
		"n.nnn       (numeric) the proof-of-work " +
		"difficulty as a multiple of the minimum " +
		"difficulty.\n" +
		"\nExamples:\n" +
		HelpExampleCli("getdifficulty") +
		HelpExampleRPC("getdifficulty")

	getmempoolancestorsDesc = "getmempoolancestors txid (verbose)\n" +
		"\nIf txid is in the mempool, returns all in-mempool ancestors.\n" +
		"\nArguments:\n" +
		"1. \"txid\"                 (string, required) The transaction id " +
		"(must be in mempool)\n" +
		"2. verbose                  (boolean, optional, default=false) " +
		"True for a json object, false for array of transaction ids\n" +
		"\nResult (for verbose=false):\n" +
		"[                       (json array of strings)\n" +
		"  \"transactionid\"           (string) The transaction id of an " +
		"in-mempool ancestor transaction\n" +
		"  ,...\n" +
		"]\n" +
		"\nResult (for verbose=true):\n" +
		"{                           (json object)\n" +
		"  \"transactionid\" : {       (json object)\n" +
		"    \"size\" : n,             (numeric) transaction size.\n" +
		"    \"fee\" : n,              (numeric) transaction fee in " +
		"BCH\n" +
		"    \"modifiedfee\" : n,      (numeric) transaction fee with fee " +
		"deltas used for mining priority\n" +
		"    \"time\" : n,             (numeric) local time transaction " +
		"entered pool in seconds since 1 Jan 1970 GMT\n" +
		"    \"height\" : n,           (numeric) block height when " +
		"transaction entered pool\n" +
		"    \"startingpriority\" : n, (numeric) DEPRECATED. Priority when " +
		"transaction entered pool\n" +
		"    \"currentpriority\" : n,  (numeric) DEPRECATED. Transaction " +
		"priority now\n" +
		"    \"descendantcount\" : n,  (numeric) number of in-mempool " +
		"descendant transactions (including this one)\n" +
		"    \"descendantsize\" : n,   (numeric) virtual transaction size " +
		"of in-mempool descendants (including this one)\n" +
		"    \"descendantfees\" : n,   (numeric) modified fees (see above) " +
		"of in-mempool descendants (including this one)\n" +
		"    \"ancestorcount\" : n,    (numeric) number of in-mempool " +
		"ancestor transactions (including this one)\n" +
		"    \"ancestorsize\" : n,     (numeric) virtual transaction size " +
		"of in-mempool ancestors (including this one)\n" +
		"    \"ancestorfees\" : n,     (numeric) modified fees (see above) " +
		"of in-mempool ancestors (including this one)\n" +
		"    \"depends\" : [           (array) unconfirmed transactions " +
		"used as inputs for this transaction\n" +
		"        \"transactionid\",    (string) parent transaction id\n" +
		"       ... ]\n" +
		"  }, ...\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getmempoolancestors", "\"mytxid\"") +
		HelpExampleRPC("getmempoolancestors", "\"mytxid\"")

	getmempooldescendantsDesc = "getmempooldescendants txid (verbose)\n" +
		"\nIf txid is in the mempool, returns all in-mempool descendants.\n" +
		"\nArguments:\n" +
		"1. \"txid\"                 (string, required) The transaction id " +
		"(must be in mempool)\n" +
		"2. verbose                  (boolean, optional, default=false) " +
		"True for a json object, false for array of transaction ids\n" +
		"\nResult (for verbose=false):\n" +
		"[                       (json array of strings)\n" +
		"  \"transactionid\"           (string) The transaction id of an " +
		"in-mempool descendant transaction\n" +
		"  ,...\n" +
		"]\n" +
		"\nResult (for verbose=true):\n" +
		"{                           (json object)\n" +
		"  \"transactionid\" : {       (json object)\n" +
		"    \"size\" : n,             (numeric) transaction size.\n" +
		"    \"fee\" : n,              (numeric) transaction fee in " +
		"BCH\n" +
		"    \"modifiedfee\" : n,      (numeric) transaction fee with fee " +
		"deltas used for mining priority\n" +
		"    \"time\" : n,             (numeric) local time transaction " +
		"entered pool in seconds since 1 Jan 1970 GMT\n" +
		"    \"height\" : n,           (numeric) block height when " +
		"transaction entered pool\n" +
		"    \"startingpriority\" : n, (numeric) DEPRECATED. Priority when " +
		"transaction entered pool\n" +
		"    \"currentpriority\" : n,  (numeric) DEPRECATED. Transaction " +
		"priority now\n" +
		"    \"descendantcount\" : n,  (numeric) number of in-mempool " +
		"descendant transactions (including this one)\n" +
		"    \"descendantsize\" : n,   (numeric) virtual transaction size " +
		"of in-mempool descendants (including this one)\n" +
		"    \"descendantfees\" : n,   (numeric) modified fees (see above) " +
		"of in-mempool descendants (including this one)\n" +
		"    \"ancestorcount\" : n,    (numeric) number of in-mempool " +
		"ancestor transactions (including this one)\n" +
		"    \"ancestorsize\" : n,     (numeric) virtual transaction size " +
		"of in-mempool ancestors (including this one)\n" +
		"    \"ancestorfees\" : n,     (numeric) modified fees (see above) " +
		"of in-mempool ancestors (including this one)\n" +
		"    \"depends\" : [           (array) unconfirmed transactions " +
		"used as inputs for this transaction\n" +
		"        \"transactionid\",    (string) parent transaction id\n" +
		"       ... ]\n" +
		"  }, ...\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getmempooldescendants", "\"mytxid\"") +
		HelpExampleRPC("getmempooldescendants", "\"mytxid\"")

	getmempoolentryDesc = "getmempoolentry txid\n" +
		"\nReturns mempool data for given transaction\n" +
		"\nArguments:\n" +
		"1. \"txid\"                   (string, required) " +
		"The transaction id (must be in mempool)\n" +
		"\nResult:\n" +
		"{                           (json object)\n" +
		"    \"size\" : n,             (numeric) transaction size.\n" +
		"    \"fee\" : n,              (numeric) transaction fee in " +
		"BCH\n" +
		"    \"modifiedfee\" : n,      (numeric) transaction fee with fee " +
		"deltas used for mining priority\n" +
		"    \"time\" : n,             (numeric) local time transaction " +
		"entered pool in seconds since 1 Jan 1970 GMT\n" +
		"    \"height\" : n,           (numeric) block height when " +
		"transaction entered pool\n" +
		"    \"startingpriority\" : n, (numeric) DEPRECATED. Priority when " +
		"transaction entered pool\n" +
		"    \"currentpriority\" : n,  (numeric) DEPRECATED. Transaction " +
		"priority now\n" +
		"    \"descendantcount\" : n,  (numeric) number of in-mempool " +
		"descendant transactions (including this one)\n" +
		"    \"descendantsize\" : n,   (numeric) virtual transaction size " +
		"of in-mempool descendants (including this one)\n" +
		"    \"descendantfees\" : n,   (numeric) modified fees (see above) " +
		"of in-mempool descendants (including this one)\n" +
		"    \"ancestorcount\" : n,    (numeric) number of in-mempool " +
		"ancestor transactions (including this one)\n" +
		"    \"ancestorsize\" : n,     (numeric) virtual transaction size " +
		"of in-mempool ancestors (including this one)\n" +
		"    \"ancestorfees\" : n,     (numeric) modified fees (see above) " +
		"of in-mempool ancestors (including this one)\n" +
		"    \"depends\" : [           (array) unconfirmed transactions " +
		"used as inputs for this transaction\n" +
		"        \"transactionid\",    (string) parent transaction id\n" +
		"       ... ]\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getmempoolentry", "\"mytxid\"") +
		HelpExampleRPC("getmempoolentry", "\"mytxid\"")

	getmempoolinfoDesc = "getmempoolinfo\n" +
		"\nReturns details on the active state of the TX memory pool.\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"size\": xxxxx,               (numeric) Current tx count\n" +
		"  \"bytes\": xxxxx,              (numeric) Transaction size.\n" +
		"  \"usage\": xxxxx,              (numeric) Total memory usage for " +
		"the mempool\n" +
		"  \"maxmempool\": xxxxx,         (numeric) Maximum memory usage " +
		"for the mempool\n" +
		"  \"mempoolminfee\": xxxxx       (numeric) Minimum fee for tx to " +
		"be accepted\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getmempoolinfo") +
		HelpExampleRPC("getmempoolinfo")

	getrawmempoolDesc = "getrawmempool ( verbose )\n" +
		"\nReturns all transaction ids in memory pool as a json array of " +
		"string transaction ids.\n" +
		"\nArguments:\n" +
		"1. verbose (boolean, optional, default=false) True for a json " +
		"object, false for array of transaction ids\n" +
		"\nResult: (for verbose = false):\n" +
		"[                     (json array of string)\n" +
		"  \"transactionid\"     (string) The transaction id\n" +
		"  ,...\n" +
		"]\n" +
		"\nResult: (for verbose = true):\n" +
		"{                           (json object)\n" +
		"  \"transactionid\" : {       (json object)\n" +
		"    \"size\" : n,             (numeric) transaction size.\n" +
		"    \"fee\" : n,              (numeric) transaction fee in " +
		"BCH\n" +
		"    \"modifiedfee\" : n,      (numeric) transaction fee with fee " +
		"deltas used for mining priority\n" +
		"    \"time\" : n,             (numeric) local time transaction " +
		"entered pool in seconds since 1 Jan 1970 GMT\n" +
		"    \"height\" : n,           (numeric) block height when " +
		"transaction entered pool\n" +
		"    \"startingpriority\" : n, (numeric) DEPRECATED. Priority when " +
		"transaction entered pool\n" +
		"    \"currentpriority\" : n,  (numeric) DEPRECATED. Transaction " +
		"priority now\n" +
		"    \"descendantcount\" : n,  (numeric) number of in-mempool " +
		"descendant transactions (including this one)\n" +
		"    \"descendantsize\" : n,   (numeric) virtual transaction size " +
		"of in-mempool descendants (including this one)\n" +
		"    \"descendantfees\" : n,   (numeric) modified fees (see above) " +
		"of in-mempool descendants (including this one)\n" +
		"    \"ancestorcount\" : n,    (numeric) number of in-mempool " +
		"ancestor transactions (including this one)\n" +
		"    \"ancestorsize\" : n,     (numeric) virtual transaction size " +
		"of in-mempool ancestors (including this one)\n" +
		"    \"ancestorfees\" : n,     (numeric) modified fees (see above) " +
		"of in-mempool ancestors (including this one)\n" +
		"    \"depends\" : [           (array) unconfirmed transactions " +
		"used as inputs for this transaction\n" +
		"        \"transactionid\",    (string) parent transaction id\n" +
		"       ... ]\n" +
		"  }, ...\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getrawmempool") +
		HelpExampleRPC("getrawmempool")

	gettxoutDesc = "gettxout \"txid\" n ( include_mempool )\n" +
		"\nReturns details about an unspent transaction output.\n" +
		"\nArguments:\n" +
		"1. \"txid\"       (string, required) The transaction id\n" +
		"2. n              (numeric, required) vout number\n" +
		"3. include_mempool  (boolean, optional) Whether to include the " +
		"mempool. Default: true. " +
		"Note that an unspent output that is spent in the mempool won't appear.\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"bestblock\" : \"hash\",    (string) the block hash\n" +
		"  \"confirmations\" : n,       (numeric) The number of " +
		"confirmations\n" +
		"  \"value\" : x.xxx,           (numeric) The transaction value " +
		"in " +
		"BCH\n" +
		"  \"scriptPubKey\" : {         (json object)\n" +
		"     \"asm\" : \"code\",       (string) \n" +
		"     \"hex\" : \"hex\",        (string) \n" +
		"     \"reqSigs\" : n,          (numeric) Number of required " +
		"signatures\n" +
		"     \"type\" : \"pubkeyhash\", (string) The type, eg pubkeyhash\n" +
		"     \"addresses\" : [          (array of string) array of " +
		"bitcoin addresses\n" +
		"        \"address\"     (string) bitcoin address\n" +
		"        ,...\n" +
		"     ]\n" +
		"  },\n" +
		"  \"coinbase\" : true|false   (boolean) Coinbase or not\n" +
		"}\n" +
		"\nExamples:\n" +
		"\nGet unspent transactions\n" +
		HelpExampleCli("listunspent") +
		"\nView the details\n" +
		HelpExampleCli("gettxout", "\"txid\"", "1") +
		"\nAs a json rpc call\n" +
		HelpExampleRPC("gettxout", "\"txid\"", "1")

	gettxoutsetinfoDesc = "gettxoutsetinfo\n" +
		"\nReturns statistics about the unspent transaction output set.\n" +
		"Note this call may take some time.\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"height\":n,     (numeric) The current block height (index)\n" +
		"  \"bestblock\": \"hex\",   (string) the best block hash hex\n" +
		"  \"transactions\": n,      (numeric) The number of transactions\n" +
		"  \"txouts\": n,            (numeric) The number of output " +
		"transactions\n" +
		"  \"bogosize\": n,          (numeric) A database-independent " +
		"metric for UTXO set size\n" +
		"  \"hash_serialized\": \"hash\",   (string) The serialized hash\n" +
		"  \"disk_size\": n,         (numeric) The estimated size of the " +
		"chainstate on disk\n" +
		"  \"total_amount\": x.xxx          (numeric) The total amount\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("gettxoutsetinfo") +
		HelpExampleRPC("gettxoutsetinfo")

	pruneblockchainDesc = "pruneblockchain\n" +
		"\nArguments:\n" +
		"1. \"height\"       (numeric, required) The block height to prune " +
		"up to. May be set to a discrete height, or a unix timestamp\n" +
		"                  to prune blocks whose block time is at least 2 " +
		"hours older than the provided timestamp.\n" +
		"\nResult:\n" +
		"n    (numeric) Height of the last block pruned.\n" +
		"\nExamples:\n" +
		HelpExampleCli("pruneblockchain", "1000") +
		HelpExampleRPC("pruneblockchain", "1000")

	verifychainDesc = "verifychain ( checklevel nblocks )\n" +
		"\nVerifies blockchain database.\n" +
		"\nArguments:\n" +
		"1. checklevel   (numeric, optional, 0-4, default=3)" +
		" How thorough the block verification is.\n" +
		"2. nblocks      (numeric, optional, default=6 , 0=all) " +
		"The number of blocks to check.\n" +
		"\nResult:\n" +
		"true|false       (boolean) Verified or not\n" +
		"\nExamples:\n" +
		HelpExampleCli("verifychain") +
		HelpExampleRPC("verifychain")

	preciousblockDesc = "preciousblock \"blockhash\"\n" +
		"\nTreats a block as if it were received before others with the " +
		"same work.\n" +
		"\nA later preciousblock call can override the effect of an " +
		"earlier one.\n" +
		"\nThe effects of preciousblock are not retained across restarts.\n" +
		"\nArguments:\n" +
		"1. \"blockhash\"   (string, required) the hash of the block to " +
		"mark as precious\n" +
		"\nResult:\n" +
		"\nExamples:\n" +
		HelpExampleCli("preciousblock", "\"blockhash\"") +
		HelpExampleRPC("preciousblock", "\"blockhash\"")

	waitforblockheightDesc = "waitforblockheight \"height\" (timeout)\n" +
		"\nWaits for (at least) block height and returns the height and " +
		"hash\n" +
		"of the current tip.\n" +
		"\nReturns the current block on timeout or exit.\n" +
		"\nArguments:\n" +
		"1. height  (required, int) Block height to wait for (int)\n" +
		"2. timeout (int, optional, default=0) Time in milliseconds to " +
		"wait for a response. 0 indicates no timeout.\n" +
		"\nResult:\n" +
		"{                           (json object)\n" +
		"  \"hash\" : {       (string) The blockhash\n" +
		"  \"height\" : {     (int) Block height\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("waitforblockheight", "\"height\"") +
		HelpExampleRPC("waitforblockheight", "\"height\"")
)

//mining
var (
	getnetworkhashpsDesc = "getnetworkhashps ( nblocks height )\n" +
		"\nReturns the estimated network hashes per second based on the " +
		"last n blocks.\n" +
		"Pass in [blocks] to override # of blocks, -1 specifies since last " +
		"difficulty change.\n" +
		"Pass in [height] to estimate the network speed at the time when a " +
		"certain block was found.\n" +
		"\nArguments:\n" +
		"1. nblocks     (numeric, optional, default=120) The number of " +
		"blocks, or -1 for blocks since last difficulty change.\n" +
		"2. height      (numeric, optional, default=-1) To estimate at the " +
		"time of the given height.\n" +
		"\nResult:\n" +
		"x             (numeric) Hashes per second estimated\n" +
		"\nExamples:\n" +
		HelpExampleCli("getnetworkhashps") +
		HelpExampleRPC("getnetworkhashps")

	getmininginfoDesc = "getmininginfo\n" +
		"\nReturns a json object containing mining-related information." +
		"\nResult:\n" +
		"{\n" +
		"  \"blocks\": nnn,             (numeric) The current block\n" +
		"  \"currentblocksize\": nnn,   (numeric) The last block size\n" +
		"  \"currentblocktx\": nnn,     (numeric) The last block " +
		"transaction\n" +
		"  \"difficulty\": xxx.xxxxx    (numeric) The current difficulty\n" +
		"  \"errors\": \"...\"            (string) Current errors\n" +
		"  \"networkhashps\": nnn,      (numeric) The network hashes per " +
		"second\n" +
		"  \"pooledtx\": n              (numeric) The size of the mempool\n" +
		"  \"chain\": \"xxxx\",           (string) current network name as " +
		"defined in BIP70 (main, test, regtest)\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getmininginfo") +
		HelpExampleRPC("getmininginfo")

	prioritisetransactionDesc = "prioritisetransaction <txid> <priority delta> <fee delta>\n" +
		"Accepts the transaction into mined blocks at a higher (or lower) " +
		"priority\n" +
		"\nArguments:\n" +
		"1. \"txid\"       (string, required) The transaction id.\n" +
		"2. priority_delta (numeric, required) The priority to add or " +
		"subtract.\n" +
		"                  The transaction selection algorithm considers " +
		"the tx as it would have a higher priority.\n" +
		"                  (priority of a transaction is calculated: " +
		"coinage * value_in_satoshis / txsize) \n" +
		"3. fee_delta      (numeric, required) The fee value (in satoshis) " +
		"to add (or subtract, if negative).\n" +
		"                  The fee is not actually paid, only the " +
		"algorithm for selecting transactions into a block\n" +
		"                  considers the transaction as it would have paid " +
		"a higher (or lower) fee.\n" +
		"\nResult:\n" +
		"true              (boolean) Returns true\n" +
		"\nExamples:\n" +
		HelpExampleCli("prioritisetransaction", "\"txid\"", "0.0", "10000") +
		HelpExampleRPC("prioritisetransaction", "\"txid\"", "0.0", "10000")

	getblocktemplateDesc = "getblocktemplate ( TemplateRequest )\n" +
		"\nIf the request parameters include a 'mode' key, that is used to " +
		"explicitly select between the default 'template' request or a " +
		"'proposal'.\n" +
		"It returns data needed to construct a block to work on.\n" +
		"For full specification, see BIPs 22, 23, 9, and 145:\n" +
		"    " +
		"https://github.com/bitcoin/bips/blob/master/bip-0022.mediawiki\n" +
		"    " +
		"https://github.com/bitcoin/bips/blob/master/bip-0023.mediawiki\n" +
		"    " +
		"https://github.com/bitcoin/bips/blob/master/" +
		"bip-0009.mediawiki#getblocktemplate_changes\n" +
		"    " +
		"https://github.com/bitcoin/bips/blob/master/bip-0145.mediawiki\n" +
		"\nArguments:\n" +
		"1. template_request         (json object, optional) A json object " +
		"in the following spec\n" +
		"     {\n" +
		"       \"mode\":\"template\"    (string, optional) This must be " +
		"set to \"template\", \"proposal\" (see BIP 23), or omitted\n" +
		"       \"capabilities\":[     (array, optional) A list of " +
		"strings\n" +
		"           \"support\"          (string) client side supported " +
		"feature, 'longpoll', 'coinbasetxn', 'coinbasevalue', 'proposal', " +
		"'serverlist', 'workid'\n" +
		"           ,...\n" +
		"       ],\n" +
		"       \"rules\":[            (array, optional) A list of " +
		"strings\n" +
		"           \"support\"          (string) client side supported " +
		"softfork deployment\n" +
		"           ,...\n" +
		"       ]\n" +
		"     }\n" +
		"\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"version\" : n,                    (numeric) The preferred " +
		"block version\n" +
		"  \"rules\" : [ \"rulename\", ... ],    (array of strings) " +
		"specific block rules that are to be enforced\n" +
		"  \"vbavailable\" : {                 (json object) set of " +
		"pending, supported versionbit (BIP 9) softfork deployments\n" +
		"      \"rulename\" : bitnumber          (numeric) identifies the " +
		"bit number as indicating acceptance and readiness for the named " +
		"softfork rule\n" +
		"      ,...\n" +
		"  },\n" +
		"  \"vbrequired\" : n,                 (numeric) bit mask of " +
		"versionbits the server requires set in submissions\n" +
		"  \"previousblockhash\" : \"xxxx\",     (string) The hash of " +
		"current highest block\n" +
		"  \"transactions\" : [                (array) contents of " +
		"non-coinbase transactions that should be included in the next " +
		"block\n" +
		"      {\n" +
		"         \"data\" : \"xxxx\",             (string) transaction " +
		"data encoded in hexadecimal (byte-for-byte)\n" +
		"         \"txid\" : \"xxxx\",             (string) transaction id " +
		"encoded in little-endian hexadecimal\n" +
		"         \"hash\" : \"xxxx\",             (string) hash encoded " +
		"in little-endian hexadecimal (including witness data)\n" +
		"         \"depends\" : [                (array) array of numbers " +
		"\n" +
		"             n                          (numeric) transactions " +
		"before this one (by 1-based index in 'transactions' list) that " +
		"must be present in the final block if this one is\n" +
		"             ,...\n" +
		"         ],\n" +
		"         \"fee\": n,                    (numeric) difference in " +
		"value between transaction inputs and outputs (in Satoshis); for " +
		"coinbase transactions, this is a negative Number of the total " +
		"collected block fees (ie, not including the block subsidy); if " +
		"key is not present, fee is unknown and clients MUST NOT assume " +
		"there isn't one\n" +
		"         \"sigops\" : n,                (numeric) total SigOps " +
		"cost, as counted for purposes of block limits; if key is not " +
		"present, sigop cost is unknown and clients MUST NOT assume it is " +
		"zero\n" +
		"         \"required\" : true|false      (boolean) if provided and " +
		"true, this transaction must be in the final block\n" +
		"      }\n" +
		"      ,...\n" +
		"  ],\n" +
		"  \"coinbaseaux\" : {                 (json object) data that " +
		"should be included in the coinbase's scriptSig content\n" +
		"      \"flags\" : \"xx\"                  (string) key name is to " +
		"be ignored, and value included in scriptSig\n" +
		"  },\n" +
		"  \"coinbasevalue\" : n,              (numeric) maximum allowable " +
		"input to coinbase transaction, including the generation award and " +
		"transaction fees (in Satoshis)\n" +
		"  \"coinbasetxn\" : { ... },          (json object) information " +
		"for coinbase transaction\n" +
		"  \"target\" : \"xxxx\",                (string) The hash target\n" +
		"  \"mintime\" : xxx,                  (numeric) The minimum " +
		"timestamp appropriate for next block time in seconds since epoch " +
		"(Jan 1 1970 GMT)\n" +
		"  \"mutable\" : [                     (array of string) list of " +
		"ways the block template may be changed \n" +
		"     \"value\"                          (string) A way the block " +
		"template may be changed, e.g. 'time', 'transactions', " +
		"'prevblock'\n" +
		"     ,...\n" +
		"  ],\n" +
		"  \"noncerange\" : \"00000000ffffffff\",(string) A range of valid " +
		"nonces\n" +
		"  \"sigoplimit\" : n,                 (numeric) limit of sigops " +
		"in blocks\n" +
		"  \"sizelimit\" : n,                  (numeric) limit of block " +
		"size\n" +
		"  \"curtime\" : ttt,                  (numeric) current timestamp " +
		"in seconds since epoch (Jan 1 1970 GMT)\n" +
		"  \"bits\" : \"xxxxxxxx\",              (string) compressed " +
		"target of next block\n" +
		"  \"height\" : n                      (numeric) The height of the " +
		"next block\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getblocktemplate") +
		HelpExampleRPC("getblocktemplate")

	submitblockDesc = "submitblock \"hexdata\" ( \"jsonparametersobject\" )\n" +
		"\nAttempts to submit new block to network.\n" +
		"The 'jsonparametersobject' parameter is currently ignored.\n" +
		"See https://en.bitcoin.it/wiki/BIP_0022 for full specification.\n" +
		"\nArguments\n" +
		"1. \"hexdata\"        (string, required) the hex-encoded block " +
		"data to submit\n" +
		"2. \"parameters\"     (string, optional) object of optional " +
		"parameters\n" +
		"    {\n" +
		"      \"workid\" : \"id\"    (string, optional) if the server " +
		"provided a workid, it MUST be included with submissions\n" +
		"    }\n" +
		"\nResult:\n" +
		"\nExamples:\n" +
		HelpExampleCli("submitblock", "\"mydata\"") +
		HelpExampleRPC("submitblock", "\"mydata\"")

	generateDesc = "generate nblocks ( maxtries )\n" +
		"\nMine up to nblocks blocks immediately (before the RPC call " +
		"returns)\n" +
		"\nArguments:\n" +
		"1. nblocks      (numeric, required) How many blocks are generated " +
		"immediately.\n" +
		"2. maxtries     (numeric, optional) How many iterations to try " +
		"(default = 1000000).\n" +
		"\nResult:\n" +
		"[ blockhashes ]     (array) hashes of blocks generated\n" +
		"\nExamples:\n" +
		"\nGenerate 11 blocks\n" +
		HelpExampleCli("generate", "11") +
		HelpExampleRPC("generate", "11")

	generatetoaddressDesc = "generatetoaddress nblocks address (maxtries)\n" +
		"\nMine blocks immediately to a specified address (before the RPC " +
		"call returns)\n" +
		"\nArguments:\n" +
		"1. nblocks      (numeric, required) How many blocks are generated " +
		"immediately.\n" +
		"2. address      (string, required) The address to send the newly " +
		"generated bitcoin to.\n" +
		"3. maxtries     (numeric, optional) How many iterations to try " +
		"(default = 1000000).\n" +
		"\nResult:\n" +
		"[ blockhashes ]     (array) hashes of blocks generated\n" +
		"\nExamples:\n" +
		HelpExampleCli("generatetoaddress", "11", "\"myaddress\"") +
		HelpExampleRPC("generatetoaddress", "11", "\"myaddress\"")
)

// net
var (
	getconnectioncountDesc = "getconnectioncount\n" +
		"\nReturns the number of connections to other nodes.\n" +
		"\nResult:\n" +
		"n          (numeric) The connection count\n" +
		"\nExamples:\n" +
		HelpExampleCli("getconnectioncount") +
		HelpExampleRPC("getconnectioncount")

	pingDesc = "ping\n" +
		"\nRequests that a ping be sent to all other nodes, to measure " +
		"ping time.\n" +
		"Results provided in getpeerinfo, pingtime and pingwait fields are " +
		"decimal seconds.\n" +
		"Ping command is handled in queue with all other commands, so it " +
		"measures processing backlog, not just network ping.\n" +
		"\nExamples:\n" +
		HelpExampleCli("ping") +
		HelpExampleRPC("ping")

	getpeerinfoDesc = "getpeerinfo\n" +
		"\nReturns data about each connected network node as a json array " +
		"of objects.\n" +
		"\nResult:\n" +
		"[\n" +
		"  {\n" +
		"    \"id\": n,                   (numeric) Peer index\n" +
		"    \"addr\":\"host:port\",      (string) The ip address and port " +
		"of the peer\n" +
		"    \"addrlocal\":\"ip:port\",   (string) local address\n" +
		"    \"services\":\"xxxxxxxxxxxxxxxx\",   (string) The services " +
		"offered\n" +
		"    \"relaytxes\":true|false,    (boolean) Whether peer has asked " +
		"us to relay transactions to it\n" +
		"    \"lastsend\": ttt,           (numeric) The time in seconds " +
		"since epoch (Jan 1 1970 GMT) of the last send\n" +
		"    \"lastrecv\": ttt,           (numeric) The time in seconds " +
		"since epoch (Jan 1 1970 GMT) of the last receive\n" +
		"    \"bytessent\": n,            (numeric) The total bytes sent\n" +
		"    \"bytesrecv\": n,            (numeric) The total bytes " +
		"received\n" +
		"    \"conntime\": ttt,           (numeric) The connection time in " +
		"seconds since epoch (Jan 1 1970 GMT)\n" +
		"    \"timeoffset\": ttt,         (numeric) The time offset in " +
		"seconds\n" +
		"    \"pingtime\": n,             (numeric) ping time (if " +
		"available)\n" +
		"    \"minping\": n,              (numeric) minimum observed ping " +
		"time (if any at all)\n" +
		"    \"pingwait\": n,             (numeric) ping wait (if " +
		"non-zero)\n" +
		"    \"version\": v,              (numeric) The peer version, such " +
		"as 7001\n" +
		"    \"subver\": \"/Satoshi:0.8.5/\",  (string) The string " +
		"version\n" +
		"    \"inbound\": true|false,     (boolean) Inbound (true) or " +
		"Outbound (false)\n" +
		"    \"addnode\": true|false,     (boolean) Whether connection was " +
		"due to addnode and is using an addnode slot\n" +
		"    \"startingheight\": n,       (numeric) The starting height " +
		"(block) of the peer\n" +
		"    \"banscore\": n,             (numeric) The ban score\n" +
		"    \"synced_headers\": n,       (numeric) The last header we " +
		"have in common with this peer\n" +
		"    \"synced_blocks\": n,        (numeric) The last block we have " +
		"in common with this peer\n" +
		"    \"inflight\": [\n" +
		"       n,                        (numeric) The heights of blocks " +
		"we're currently asking from this peer\n" +
		"       ...\n" +
		"    ],\n" +
		"    \"whitelisted\": true|false, (boolean) Whether the peer is " +
		"whitelisted\n" +
		"    \"bytessent_per_msg\": {\n" +
		"       \"addr\": n,              (numeric) The total bytes sent " +
		"aggregated by message type\n" +
		"       ...\n" +
		"    },\n" +
		"    \"bytesrecv_per_msg\": {\n" +
		"       \"addr\": n,              (numeric) The total bytes " +
		"received aggregated by message type\n" +
		"       ...\n" +
		"    }\n" +
		"  }\n" +
		"  ,...\n" +
		"]\n" +
		"\nExamples:\n" +
		HelpExampleCli("getpeerinfo") +
		HelpExampleRPC("getpeerinfo")

	addnodeDesc = "addnode \"node\" \"add|remove|onetry\"\n" +
		"\nAttempts add or remove a node from the addnode list.\n" +
		"Or try a connection to a node once.\n" +
		"\nArguments:\n" +
		"1. \"node\"     (string, required) The node (see getpeerinfo for " +
		"nodes)\n" +
		"2. \"command\"  (string, required) 'add' to add a node to the " +
		"list, 'remove' to remove a node from the list, 'onetry' to try a " +
		"connection to the node once\n" +
		"\nExamples:\n" +
		HelpExampleCli("addnode", "\"192.168.0.6:8333\"", "\"onetry\"") +
		HelpExampleRPC("addnode", "\"192.168.0.6:8333\"", "\"onetry\"")

	disconnectnodeDesc = "disconnectnode \"[address]\" [nodeid]\n" +
		"\nImmediately disconnects from the specified peer node.\n" +
		"\nStrictly one out of 'address' and 'nodeid' can be provided to " +
		"identify the node.\n" +
		"\nTo disconnect by nodeid, either set 'address' to the empty " +
		"string, or call using the named 'nodeid' argument only.\n" +
		"\nArguments:\n" +
		"1. \"address\"     (string, optional) The IP address/port of the " +
		"node\n" +
		"2. \"nodeid\"      (number, optional) The node ID (see " +
		"getpeerinfo for node IDs)\n" +
		"\nExamples:\n" +
		HelpExampleCli("disconnectnode", "\"192.168.0.6:8333\"") +
		HelpExampleCli("disconnectnode", "\"\"", "1") +
		HelpExampleRPC("disconnectnode", "\"192.168.0.6:8333\"") +
		HelpExampleRPC("disconnectnode", "\"\"", "1")

	getaddednodeinfoDesc = "getaddednodeinfo ( \"node\" )\n" +
		"\nReturns information about the given added node, or all added " +
		"nodes\n" +
		"(note that onetry addnodes are not listed here)\n" +
		"\nArguments:\n" +
		"1. \"node\"   (string, optional) If provided, return information " +
		"about this specific node, otherwise all nodes are returned.\n" +
		"\nResult:\n" +
		"[\n" +
		"  {\n" +
		"    \"addednode\" : \"192.168.0.201\",   (string) The node ip " +
		"address or name (as provided to addnode)\n" +
		"    \"connected\" : true|false,          (boolean) If connected\n" +
		"    \"addresses\" : [                    (list of objects) Only " +
		"when connected = true\n" +
		"       {\n" +
		"         \"address\" : \"192.168.0.201:8333\",  (string) The " +
		"bitcoin server IP and port we're connected to\n" +
		"         \"connected\" : \"outbound\"           (string) " +
		"connection, inbound or outbound\n" +
		"       }\n" +
		"     ]\n" +
		"  }\n" +
		"  ,...\n" +
		"]\n" +
		"\nExamples:\n" +
		HelpExampleCli("getaddednodeinfo", "\"192.168.0.201\"") +
		HelpExampleRPC("getaddednodeinfo", "\"192.168.0.201\"")

	getnettotalsDesc = "getnettotals\n" +
		"\nReturns information about network traffic, including bytes in, " +
		"bytes out,\n" +
		"and current time.\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"totalbytesrecv\": n,   (numeric) Total bytes received\n" +
		"  \"totalbytessent\": n,   (numeric) Total bytes sent\n" +
		"  \"timemillis\": t,       (numeric) Current UNIX time in " +
		"milliseconds\n" +
		"  \"uploadtarget\":\n" +
		"  {\n" +
		"    \"timeframe\": n,                         (numeric) Length of " +
		"the measuring timeframe in seconds\n" +
		"    \"target\": n,                            (numeric) Target in " +
		"bytes\n" +
		"    \"target_reached\": true|false,           (boolean) True if " +
		"target is reached\n" +
		"    \"serve_historical_blocks\": true|false,  (boolean) True if " +
		"serving historical blocks\n" +
		"    \"bytes_left_in_cycle\": t,               (numeric) Bytes " +
		"left in current time cycle\n" +
		"    \"time_left_in_cycle\": t                 (numeric) Seconds " +
		"left in current time cycle\n" +
		"  }\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getnettotals") +
		HelpExampleRPC("getnettotals")

	getnetworkinfoDesc = "getnetworkinfo\n" +
		"Returns an object containing various state info regarding P2P " +
		"networking.\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"version\": xxxxx,                      (numeric) the server " +
		"version\n" +
		"  \"subversion\": \"/Satoshi:x.x.x/\",     (string) the server " +
		"subversion string\n" +
		"  \"protocolversion\": xxxxx,              (numeric) the protocol " +
		"version\n" +
		"  \"localservices\": \"xxxxxxxxxxxxxxxx\", (string) the services " +
		"we offer to the network\n" +
		"  \"localrelay\": true|false,              (bool) true if " +
		"transaction relay is requested from peers\n" +
		"  \"timeoffset\": xxxxx,                   (numeric) the time " +
		"offset\n" +
		"  \"connections\": xxxxx,                  (numeric) the number " +
		"of connections\n" +
		"  \"networkactive\": true|false,           (bool) whether p2p " +
		"networking is enabled\n" +
		"  \"networks\": [                          (array) information " +
		"per network\n" +
		"  {\n" +
		"    \"name\": \"xxx\",                     (string) network " +
		"(ipv4, ipv6 or onion)\n" +
		"    \"limited\": true|false,               (boolean) is the " +
		"network limited using -onlynet?\n" +
		"    \"reachable\": true|false,             (boolean) is the " +
		"network reachable?\n" +
		"    \"proxy\": \"host:port\"               (string) the proxy " +
		"that is used for this network, or empty if none\n" +
		"    \"proxy_randomize_credentials\": true|false,  (string) " +
		"Whether randomized credentials are used\n" +
		"  }\n" +
		"  ,...\n" +
		"  ],\n" +
		"  \"relayfee\": x.xxxxxxxx,                (numeric) minimum " +
		"relay fee for non-free transactions in " +
		"BCH/kB\n" +
		"  \"incrementalfee\": x.xxxxxxxx,          " +
		"(numeric) minimum fee increment for mempool " +
		"limiting or BIP 125 replacement in BCH/kB\n" +
		"  \"localaddresses\": [                    " +
		"(array) list of local addresses\n" +
		"  {\n" +
		"    \"address\": \"xxxx\",                 " +
		"(string) network address\n" +
		"    \"port\": xxx,                         " +
		"(numeric) network port\n" +
		"    \"score\": xxx                         " +
		"(numeric) relative score\n" +
		"  }\n" +
		"  ,...\n" +
		"  ]\n" +
		"  \"warnings\": \"...\"                    " +
		"(string) any network warnings\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getnetworkinfo") +
		HelpExampleRPC("getnetworkinfo")

	setbanDesc = "setban \"subnet\" \"add|remove\" (bantime) (absolute)\n" +
		"\nAttempts add or remove a IP/Subnet from the banned list.\n" +
		"\nArguments:\n" +
		"1. \"subnet\"       (string, required) The IP/Subnet (see " +
		"getpeerinfo for nodes ip) with a optional netmask (default is /32 " +
		"= single ip)\n" +
		"2. \"command\"      (string, required) 'add' to add a IP/Subnet " +
		"to the list, 'remove' to remove a IP/Subnet from the list\n" +
		"3. \"bantime\"      (numeric, optional) time in seconds how long " +
		"(or until when if [absolute] is set) the ip is banned (0 or empty " +
		"means using the default time of 24h which can also be overwritten " +
		"by the -bantime startup argument)\n" +
		"4. \"absolute\"     (boolean, optional) If set, the bantime must " +
		"be a absolute timestamp in seconds since epoch (Jan 1 1970 GMT)\n" +
		"\nExamples:\n" +
		HelpExampleCli("setban", "\"192.168.0.6\"", "\"add\"", "86400") +
		HelpExampleCli("setban", "\"192.168.0.0/24\"", "\"add\"") +
		HelpExampleRPC("setban", "\"192.168.0.6\"", "\"add\"", "86400")

	listbannedDesc = "listbanned\n" +
		"\nList all banned IPs/Subnets.\n" +
		"\nExamples:\n" +
		HelpExampleCli("listbanned") +
		HelpExampleRPC("listbanned")

	clearbannedDesc = "clearbanned\n" +
		"\nClear all banned IPs.\n" +
		"\nExamples:\n" +
		HelpExampleCli("clearbanned") +
		HelpExampleRPC("clearbanned")

	setnetworkactiveDesc = "setnetworkactive true|false\n" +
		"\nDisable/enable all p2p network activity.\n" +
		"\nArguments:\n" +
		"1. \"state\"        (boolean, required) true to " +
		"enable networking, false to disable\n"
)

// rawtransaction
var (
	getrawtransactionDesc = "getrawtransaction \"txid\" ( verbose )\n" +
		"\nNOTE: By default this function only works for mempool " +
		"transactions. If the -txindex option is\n" +
		"enabled, it also works for blockchain transactions.\n" +
		"DEPRECATED: for now, it also works for transactions with unspent " +
		"outputs.\n" +
		"\nReturn the raw transaction data.\n" +
		"\nIf verbose is 'true', returns an Object with information about " +
		"'txid'.\n" +
		"If verbose is 'false' or omitted, returns a string that is " +
		"serialized, hex-encoded data for 'txid'.\n" +
		"\nArguments:\n" +
		"1. \"txid\"      (string, required) The transaction id\n" +
		"2. verbose       (bool, optional, default=false) If false, return " +
		"a string, otherwise return a json object\n" +
		"\nResult (if verbose is not set or set to false):\n" +
		"\"data\"      (string) The serialized, hex-encoded data for " +
		"'txid'\n" +
		"\nResult (if verbose is set to true):\n" +
		"{\n" +
		"  \"hex\" : \"data\",       (string) The serialized, hex-encoded " +
		"data for 'txid'\n" +
		"  \"txid\" : \"id\",        (string) The transaction id (same as " +
		"provided)\n" +
		"  \"hash\" : \"id\",        (string) The transaction hash " +
		"(differs from txid for witness transactions)\n" +
		"  \"size\" : n,             (numeric) The serialized transaction " +
		"size\n" +
		"  \"version\" : n,          (numeric) The version\n" +
		"  \"locktime\" : ttt,       (numeric) The lock time\n" +
		"  \"vin\" : [               (array of json objects)\n" +
		"     {\n" +
		"       \"txid\": \"id\",    (string) The transaction id\n" +
		"       \"vout\": n,         (numeric) \n" +
		"       \"scriptSig\": {     (json object) The script\n" +
		"         \"asm\": \"asm\",  (string) asm\n" +
		"         \"hex\": \"hex\"   (string) hex\n" +
		"       },\n" +
		"       \"sequence\": n      (numeric) The script sequence number\n" +
		"     }\n" +
		"     ,...\n" +
		"  ],\n" +
		"  \"vout\" : [              (array of json objects)\n" +
		"     {\n" +
		"       \"value\" : x.xxx,            (numeric) The value in BCH\n" +
		"       \"n\" : n,                    (numeric) index\n" +
		"       \"scriptPubKey\" : {          (json object)\n" +
		"         \"asm\" : \"asm\",          (string) the asm\n" +
		"         \"hex\" : \"hex\",          (string) the hex\n" +
		"         \"reqSigs\" : n,            (numeric) The required sigs\n" +
		"         \"type\" : \"pubkeyhash\",  (string) The type, eg " +
		"'pubkeyhash'\n" +
		"         \"addresses\" : [           (json array of string)\n" +
		"           \"address\"        (string) bitcoin address\n" +
		"           ,...\n" +
		"         ]\n" +
		"       }\n" +
		"     }\n" +
		"     ,...\n" +
		"  ],\n" +
		"  \"blockhash\" : \"hash\",   (string) the block hash\n" +
		"  \"confirmations\" : n,      (numeric) The confirmations\n" +
		"  \"time\" : ttt,             (numeric) The transaction time in " +
		"seconds since epoch (Jan 1 1970 GMT)\n" +
		"  \"blocktime\" : ttt         (numeric) The block time in seconds " +
		"since epoch (Jan 1 1970 GMT)\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getrawtransaction", "\"mytxid\"") +
		HelpExampleCli("getrawtransaction", "\"mytxid\"", "true") +
		HelpExampleRPC("getrawtransaction", "\"mytxid\"", "true")

	createrawtransactionDesc = "createrawtransaction [{\"txid\":\"id\",\"vout\":n},...] " +
		"{\"address\":amount,\"data\":\"hex\",...} ( locktime )\n" +
		"\nCreate a transaction spending the given inputs and creating new " +
		"outputs.\n" +
		"Outputs can be addresses or data.\n" +
		"Returns hex-encoded raw transaction.\n" +
		"Note that the transaction's inputs are not signed, and\n" +
		"it is not stored in the wallet or transmitted to the network.\n" +
		"\nArguments:\n" +
		"1. \"inputs\"                (array, required) A json array of " +
		"json objects\n" +
		"     [\n" +
		"       {\n" +
		"         \"txid\":\"id\",    (string, required) The transaction " +
		"id\n" +
		"         \"vout\":n,         (numeric, required) The output " +
		"number\n" +
		"         \"sequence\":n      (numeric, optional) The sequence " +
		"number\n" +
		"       } \n" +
		"       ,...\n" +
		"     ]\n" +
		"2. \"outputs\"               (object, required) a json object " +
		"with outputs\n" +
		"    {\n" +
		"      \"address\": x.xxx,    (numeric or string, required) The " +
		"key is the bitcoin address, the numeric value (can be string) is " +
		"the BCH" +
		" amount\n" +
		"      \"data\": \"hex\"      (string, required) The key is " +
		"\"data\", the value is hex encoded data\n" +
		"      ,...\n" +
		"    }\n" +
		"3. locktime                  (numeric, optional, default=0) Raw " +
		"locktime. Non-0 value also locktime-activates inputs\n" +
		"\nResult:\n" +
		"\"transaction\"              (string) hex string of the " +
		"transaction\n" +
		"\nExamples:\n" +
		HelpExampleCli("createrawtransaction", `"[{\"txid\":\"myid\",\"vout\":0}]"`, `"{\"address\":0.01}"`) +
		HelpExampleCli("createrawtransaction", `"[{\"txid\":\"myid\",\"vout\":0}]"`, `"{\"data\":\"00010203\"}"`) +
		HelpExampleRPC("createrawtransaction", `"[{\"txid\":\"myid\",\"vout\":0}]"`, `"{\"address\":0.01}"`) +
		HelpExampleRPC("createrawtransaction", `"[{\"txid\":\"myid\",\"vout\":0}]"`, `"{\"data\":\"00010203\"}"`)

	decoderawtransactionDesc = "decoderawtransaction \"hexstring\"\n" +
		"\nReturn a JSON object representing the serialized, hex-encoded " +
		"transaction.\n" +
		"\nArguments:\n" +
		"1. \"hexstring\"      (string, required) The transaction hex " +
		"string\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"txid\" : \"id\",        (string) The transaction id\n" +
		"  \"hash\" : \"id\",        (string) The transaction hash " +
		"(differs from txid for witness transactions)\n" +
		"  \"size\" : n,             (numeric) The transaction size\n" +
		"  \"version\" : n,          (numeric) The version\n" +
		"  \"locktime\" : ttt,       (numeric) The lock time\n" +
		"  \"vin\" : [               (array of json objects)\n" +
		"     {\n" +
		"       \"txid\": \"id\",    (string) The transaction id\n" +
		"       \"vout\": n,         (numeric) The output number\n" +
		"       \"scriptSig\": {     (json object) The script\n" +
		"         \"asm\": \"asm\",  (string) asm\n" +
		"         \"hex\": \"hex\"   (string) hex\n" +
		"       },\n" +
		"       \"sequence\": n     (numeric) The script sequence number\n" +
		"     }\n" +
		"     ,...\n" +
		"  ],\n" +
		"  \"vout\" : [             (array of json objects)\n" +
		"     {\n" +
		"       \"value\" : x.xxx,            (numeric) The value in BCH\n" +
		"       \"n\" : n,                    (numeric) index\n" +
		"       \"scriptPubKey\" : {          (json object)\n" +
		"         \"asm\" : \"asm\",          (string) the asm\n" +
		"         \"hex\" : \"hex\",          (string) the hex\n" +
		"         \"reqSigs\" : n,            (numeric) The required sigs\n" +
		"         \"type\" : \"pubkeyhash\",  (string) The type, eg " +
		"'pubkeyhash'\n" +
		"         \"addresses\" : [           (json array of string)\n" +
		"           \"12tvKAXCxZjSmdNbao16dKXC8tRWfcF5oc\"   (string) " +
		"bitcoin address\n" +
		"           ,...\n" +
		"         ]\n" +
		"       }\n" +
		"     }\n" +
		"     ,...\n" +
		"  ],\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("decoderawtransaction", "\"hexstring\"") +
		HelpExampleRPC("decoderawtransaction", "\"hexstring\"")

	decodescriptDesc = "decodescript \"hexstring\"\n" +
		"\nDecode a hex-encoded script.\n" +
		"\nArguments:\n" +
		"1. \"hexstring\"     (string) the hex encoded script\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"asm\":\"asm\",   (string) Script public key\n" +
		"  \"hex\":\"hex\",   (string) hex encoded public key\n" +
		"  \"type\":\"type\", (string) The output type\n" +
		"  \"reqSigs\": n,    (numeric) The required signatures\n" +
		"  \"addresses\": [   (json array of string)\n" +
		"     \"address\"     (string) bitcoin address\n" +
		"     ,...\n" +
		"  ],\n" +
		"  \"p2sh\",\"address\" (string) address of P2SH script wrapping " +
		"this redeem script (not returned if the script is already a " +
		"P2SH).\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("decodescript", "\"hexstring\"") +
		HelpExampleRPC("decodescript", "\"hexstring\"")

	sendrawtransactionDesc = "sendrawtransaction \"hexstring\" ( allowhighfees )\n" +
		"\nSubmits raw transaction (serialized, hex-encoded) to local node " +
		"and network.\n" +
		"\nAlso see createrawtransaction and signrawtransaction calls.\n" +
		"\nArguments:\n" +
		"1. \"hexstring\"    (string, required) The hex string of the raw " +
		"transaction)\n" +
		"2. allowhighfees    (boolean, optional, default=false) Allow high " +
		"fees\n" +
		"\nResult:\n" +
		"\"hex\"             (string) The transaction hash in hex\n" +
		"\nExamples:\n" +
		"\nCreate a transaction\n" +
		HelpExampleCli("createrawtransaction",
			"\"[{\\\"txid\\\" : "+
				"\\\"mytxid\\\",\\\"vout\\\":0}]\" "+
				"\"{\\\"myaddress\\\":0.01}\"") +
		"Sign the transaction, and get back the hex\n" +
		HelpExampleCli("signrawtransaction", "\"myhex\"") +
		"\nSend the transaction (signed hex)\n" +
		HelpExampleCli("sendrawtransaction", "\"signedhex\"") +
		"\nAs a json rpc call\n" +
		HelpExampleRPC("sendrawtransaction", "\"signedhex\"")

	signrawtransactionDesc = "signrawtransaction \"hexstring\" ( " +
		"[{\"txid\":\"id\",\"vout\":n,\"scriptPubKey\":\"hex\"," +
		"\"redeemScript\":\"hex\"},...] [\"privatekey1\",...] sighashtype " +
		")\n" +
		"\nSign inputs for raw transaction (serialized, hex-encoded).\n" +
		"The second optional argument (may be null) is an array of " +
		"previous transaction outputs that\n" +
		"this transaction depends on but may not yet be in the block " +
		"chain.\n" +
		"The third optional argument (may be null) is an array of " +
		"base58-encoded private\n" +
		"keys that, if given, will be the only keys used to sign the " +
		"transaction.\n" +
		"\nArguments:\n" +
		"1. \"hexstring\"     (string, required) The transaction hex " +
		"string\n" +
		"2. \"prevtxs\"       (string, optional) An json array of previous " +
		"dependent transaction outputs\n" +
		"     [               (json array of json objects, or 'null' if " +
		"none provided)\n" +
		"       {\n" +
		"         \"txid\":\"id\",             (string, required) The " +
		"transaction id\n" +
		"         \"vout\":n,                  (numeric, required) The " +
		"output number\n" +
		"         \"scriptPubKey\": \"hex\",   (string, required) script " +
		"key\n" +
		"         \"redeemScript\": \"hex\",   (string, required for P2SH " +
		"or P2WSH) redeem script\n" +
		"         \"amount\": value            (numeric, required) The " +
		"amount spent\n" +
		"       }\n" +
		"       ,...\n" +
		"    ]\n" +
		"3. \"privkeys\"     (string, optional) A json array of " +
		"base58-encoded private keys for signing\n" +
		"    [                  (json array of strings, or 'null' if none " +
		"provided)\n" +
		"      \"privatekey\"   (string) private key in base58-encoding\n" +
		"      ,...\n" +
		"    ]\n" +
		"4. \"sighashtype\"     (string, optional, default=ALL) The " +
		"signature hash type. Must be one of\n" +
		"       \"ALL\"\n" +
		"       \"NONE\"\n" +
		"       \"SINGLE\"\n" +
		"       \"ALL|ANYONECANPAY\"\n" +
		"       \"NONE|ANYONECANPAY\"\n" +
		"       \"SINGLE|ANYONECANPAY\"\n" +
		"       \"ALL|FORKID\"\n" +
		"       \"NONE|FORKID\"\n" +
		"       \"SINGLE|FORKID\"\n" +
		"       \"ALL|FORKID|ANYONECANPAY\"\n" +
		"       \"NONE|FORKID|ANYONECANPAY\"\n" +
		"       \"SINGLE|FORKID|ANYONECANPAY\"\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"hex\" : \"value\",           (string) The hex-encoded raw " +
		"transaction with signature(s)\n" +
		"  \"complete\" : true|false,   (boolean) If the transaction has a " +
		"complete set of signatures\n" +
		"  \"errors\" : [                 (json array of objects) Script " +
		"verification errors (if there are any)\n" +
		"    {\n" +
		"      \"txid\" : \"hash\",           (string) The hash of the " +
		"referenced, previous transaction\n" +
		"      \"vout\" : n,                (numeric) The index of the " +
		"output to spent and used as input\n" +
		"      \"scriptSig\" : \"hex\",       (string) The hex-encoded " +
		"signature script\n" +
		"      \"sequence\" : n,            (numeric) Script sequence " +
		"number\n" +
		"      \"error\" : \"text\"           (string) Verification or " +
		"signing error related to the input\n" +
		"    }\n" +
		"    ,...\n" +
		"  ]\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("signrawtransaction", "\"myhex\"") +
		HelpExampleRPC("signrawtransaction", "\"myhex\"")

	gettxoutproofDesc = "gettxoutproof [\"txid\",...] ( blockhash )\n" +
		"\nReturns a hex-encoded proof that \"txid\" was included in a " +
		"block.\n" +
		"\nNOTE: By default this function only works sometimes. This is " +
		"when there is an\n" +
		"unspent output in the utxo for this transaction. To make it " +
		"always work,\n" +
		"you need to maintain a transaction index, using the -txindex " +
		"command line option or\n" +
		"specify the block in which the transaction is included manually " +
		"(by blockhash).\n" +
		"\nArguments:\n" +
		"1. \"txids\"       (string) A json array of txids to filter\n" +
		"    [\n" +
		"      \"txid\"     (string) A transaction hash\n" +
		"      ,...\n" +
		"    ]\n" +
		"2. \"blockhash\"   (string, optional) If specified, looks for " +
		"txid in the block with this hash\n" +
		"\nResult:\n" +
		"\"data\"           (string) A string that is a serialized, " +
		"hex-encoded data for the proof.\n"

	verifytxoutproofDesc = "verifytxoutproof \"proof\"\n" +
		"\nVerifies that a proof points to a transaction in a block, " +
		"returning the transaction it commits to\n" +
		"and throwing an RPC error if the block is not in our best chain\n" +
		"\nArguments:\n" +
		"1. \"proof\"    (string, required) The hex-encoded proof " +
		"generated by gettxoutproof\n" +
		"\nResult:\n" +
		"[\"txid\"]      (array, strings) The txid(s) which the proof " +
		"commits to, or empty array if the proof is invalid\n"
)

// misc
var (
	getinfoDesc = "getinfo\n" +
		"\nDEPRECATED. Returns an object containing various state info.\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"version\": xxxxx,           (numeric) the server version\n" +
		"  \"protocolversion\": xxxxx,   (numeric) the protocol version\n" +
		"  \"walletversion\": xxxxx,     (numeric) the wallet version\n" +
		"  \"balance\": xxxxxxx,         (numeric) the total bitcoin " +
		"balance of the wallet\n" +
		"  \"blocks\": xxxxxx,           (numeric) the current number of " +
		"blocks processed in the server\n" +
		"  \"timeoffset\": xxxxx,        (numeric) the time offset\n" +
		"  \"connections\": xxxxx,       (numeric) the number of " +
		"connections\n" +
		"  \"proxy\": \"host:port\",     (string, optional) the proxy used " +
		"by the server\n" +
		"  \"difficulty\": xxxxxx,       (numeric) the current difficulty\n" +
		"  \"testnet\": true|false,      (boolean) if the server is using " +
		"testnet or not\n" +
		"  \"keypoololdest\": xxxxxx,    (numeric) the timestamp (seconds " +
		"since Unix epoch) of the oldest pre-generated key in the key " +
		"pool\n" +
		"  \"keypoolsize\": xxxx,        (numeric) how many new keys are " +
		"pre-generated\n" +
		"  \"unlocked_until\": ttt,      (numeric) the timestamp in " +
		"seconds since epoch (midnight Jan 1 1970 GMT) that the wallet is " +
		"unlocked for transfers, or 0 if the wallet is locked\n" +
		"  \"paytxfee\": x.xxxx,         (numeric) the transaction fee set in BCH/kB" +
		"\"relayfee\": x.xxxx,           (numeric) minimum relay fee for non-free transactions in BCH/kB" +
		"  \"errors\": \"...\"           (string) any error messages\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("getinfo") +
		HelpExampleRPC("getinfo")

	helpDesc = "help ( \"command\" )\n" +
		"\nList all commands, or get help for a specified command.\n" +
		"\nArguments:\n" +
		"1. \"command\"     (string, optional) The command to get help on\n" +
		"\nResult:\n" +
		"\"text\"     (string) The help text\n"

	stopDesc = "stop\n" +
		"\nStop Copernicus server."

	validateaddressDesc = "validateaddress \"address\"\n" +
		"\nReturn information about the given bitcoin address.\n" +
		"\nArguments:\n" +
		"1. \"address\"     (string, required) The bitcoin address to " +
		"validate\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"isvalid\" : true|false,       (boolean) If the address is " +
		"valid or not. If not, this is the only property returned.\n" +
		"  \"address\" : \"address\", (string) The bitcoin address " +
		"validated\n" +
		"  \"scriptPubKey\" : \"hex\",       (string) The hex encoded " +
		"scriptPubKey generated by the address\n" +
		"  \"ismine\" : true|false,        (boolean) If the address is " +
		"yours or not\n" +
		"  \"iswatchonly\" : true|false,   (boolean) If the address is " +
		"watchonly\n" +
		"  \"isscript\" : true|false,      (boolean) If the key is a " +
		"script\n" +
		"  \"pubkey\" : \"publickeyhex\",    (string) The hex value of the " +
		"raw public key\n" +
		"  \"iscompressed\" : true|false,  (boolean) If the address is " +
		"compressed\n" +
		"  \"account\" : \"account\"         (string) DEPRECATED. The " +
		"account associated with the address, \"\" is the default account\n" +
		"  \"timestamp\" : timestamp,        (number, optional) The " +
		"creation time of the key if available in seconds since epoch (Jan " +
		"1 1970 GMT)\n" +
		"  \"hdkeypath\" : \"keypath\"       (string, optional) The HD " +
		"keypath if the key is HD and available\n" +
		"  \"hdmasterkeyid\" : \"<hash160>\" (string, optional) The " +
		"Hash160 of the HD master pubkey\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("validateaddress", "\"1PSSGeFHDnKNxiEyFrD1wcEaHr9hrQDDWc\"") +
		HelpExampleRPC("validateaddress", "\"1PSSGeFHDnKNxiEyFrD1wcEaHr9hrQDDWc\"")

	createmultisigDesc = "createmultisig nrequired [\"key\",...]\n" +
		"\nCreates a multi-signature address with n signature of m keys " +
		"required.\n" +
		"It returns a json object with the address and redeemScript.\n" +
		"\nArguments:\n" +
		"1. nrequired      (numeric, required) The number of required " +
		"signatures out of the n keys or addresses.\n" +
		"2. \"keys\"       (string, required) A json array of keys which " +
		"are bitcoin addresses or hex-encoded public keys\n" +
		"     [\n" +
		"       \"key\"    (string) bitcoin address or hex-encoded public " +
		"key\n" +
		"       ,...\n" +
		"     ]\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"address\":\"multisigaddress\",  (string) The value of the new " +
		"multisig address.\n" +
		"  \"redeemScript\":\"script\"       (string) The string value of " +
		"the hex-encoded redemption script.\n" +
		"}\n" +
		"\nExamples:\n" +
		"\nCreate a multisig address from 2 addresses\n" +
		HelpExampleCli("createmultisig", "2",
			"\"[\\\"16sSauSf5pF2UkUwvKGq4qjNRzBZYqgEL5\\\",\\\"171sgjn4YtPu27adkKGrdDwzRTxnRkBfKV\\\"]\"") +
		"\nAs a json rpc call\n" +
		HelpExampleRPC("createmultisig", "2",
			"\"[\\\"16sSauSf5pF2UkUwvKGq4qjNRzBZYqgEL5\\\",\\\"171sgjn4YtPu27adkKGrdDwzRTxnRkBfKV\\\"]\"")

	echoDesc = "echo \"message\" ...\n" +
		"\nSimply echo back the input arguments. This command is for testing."

	uptimeDesc = "uptime\n" +
		"\nReturns the total uptime of the server.\n" +
		"\nResult:\n" +
		"ttt        (numeric) The number of seconds " +
		"that the server has been running\n" +
		"\nExamples:\n" +
		HelpExampleCli("uptime") +
		HelpExampleRPC("uptime")
)

// wallet
var (
	getnewaddressDesc = "getnewaddress ( \"account\" )\n" +
		"\nReturns a new Bitcoin address for receiving payments.\n" +
		"If 'account' is specified (DEPRECATED), it is added to the " +
		"address book \n" +
		"so payments received with the address will be credited to " +
		"'account'.\n" +
		"\nArguments:\n" +
		"1. \"account\"        (string, optional) DEPRECATED. The account " +
		"name for the address to be linked to. If not provided, the " +
		"default account \"\" is used. It can also be set to the empty " +
		"string \"\" to represent the default account. The account does " +
		"not need to exist, it will be created if there is no account by " +
		"the given name.\n" +
		"\nResult:\n" +
		"\"address\"    (string) The new bitcoin address\n" +
		"\nExamples:\n" +
		HelpExampleCli("getnewaddress") +
		HelpExampleRPC("getnewaddress")

	listunspentDesc = "listunspent ( minconf maxconf  [\"addresses\",...] " +
		"[include_unsafe] )\n" +
		"\nReturns array of unspent transaction outputs\n" +
		"with between minconf and maxconf (inclusive) confirmations.\n" +
		"Optionally filter to only include txouts paid to specified " +
		"addresses.\n" +
		"\nArguments:\n" +
		"1. minconf          (numeric, optional, default=1) The minimum " +
		"confirmations to filter\n" +
		"2. maxconf          (numeric, optional, default=9999999) The " +
		"maximum confirmations to filter\n" +
		"3. \"addresses\"    (string) A json array of bitcoin addresses to " +
		"filter\n" +
		"    [\n" +
		"      \"address\"   (string) bitcoin address\n" +
		"      ,...\n" +
		"    ]\n" +
		"4. include_unsafe (bool, optional, default=true) Include outputs " +
		"that are not safe to spend\n" +
		"                  because they come from unconfirmed untrusted " +
		"transactions or unconfirmed\n" +
		"                  replacement transactions (cases where we are " +
		"less sure that a conflicting\n" +
		"                  transaction won't be mined).\n" +
		"\nResult\n" +
		"[                   (array of json object)\n" +
		"  {\n" +
		"    \"txid\" : \"txid\",          (string) the transaction id \n" +
		"    \"vout\" : n,               (numeric) the vout value\n" +
		"    \"address\" : \"address\",    (string) the bitcoin address\n" +
		"    \"account\" : \"account\",    (string) DEPRECATED. The " +
		"associated account, or \"\" for the default account\n" +
		"    \"scriptPubKey\" : \"key\",   (string) the script key\n" +
		"    \"amount\" : x.xxx,         (numeric) the transaction output " +
		"amount in " +
		util.CurrencyUnit +
		"\n" +
		"    \"confirmations\" : n,      (numeric) The number of " +
		"confirmations\n" +
		"    \"redeemScript\" : \"hex\",   (string) The redeemScript if " +
		"scriptPubKey is P2SH\n" +
		"    \"spendable\" : xxx,        (bool) Whether we have the " +
		"private keys to spend this output\n" +
		"    \"solvable\" : xxx,         (bool) Whether we know how to " +
		"spend this output, ignoring the lack of keys\n" +
		"    \"safe\" : xxx              (bool) Whether this output is " +
		"considered safe to spend. Unconfirmed transactions\n" +
		"                              from outside keys are considered " +
		"unsafe and are not eligible for spending by\n" +
		"                              fundrawtransaction and " +
		"sendtoaddress.\n" +
		"  }\n" +
		"  ,...\n" +
		"]\n" +
		"\nExamples\n" +
		HelpExampleCli("listunspent") +
		HelpExampleCli("listunspent", "6", "9999999",
			"\"[\\\"1PGFqEzfmQch1gKD3ra4k18PNj3tTUUSqg\\\",\\\"1LtvqCaApEdUGFkpKMM4MstjcaL4dKg8SP\\\"]\"") +
		HelpExampleRPC("listunspent", "6", "9999999",
			"\"[\\\"1PGFqEzfmQch1gKD3ra4k18PNj3tTUUSqg\\\",\\\"1LtvqCaApEdUGFkpKMM4MstjcaL4dKg8SP\\\"]\"") +
		HelpExampleCli("listunspent", "6", "9999999",
			"'[]'", "true", "'{ \"minimumAmount\": 0.005 }'") +
		HelpExampleRPC("listunspent", "6", "9999999",
			"[]", "true", "'{ \"minimumAmount\": 0.005 }'")

	settxfeeDesc = "settxfee amount\n" +
		"\nSet the transaction fee per kB. Overwrites the paytxfee " +
		"parameter.\n" +
		"\nArguments:\n" +
		"1. amount         (numeric or string, required) The transaction " +
		"fee in " +
		util.CurrencyUnit +
		"/kB\n" +
		"\nResult\n" +
		"true|false        (boolean) Returns true if successful\n" +
		"\nExamples:\n" +
		HelpExampleCli("settxfee", "0.00001") +
		HelpExampleRPC("settxfee", "0.00001")

	sendtoaddressDesc = "sendtoaddress \"address\" amount ( \"comment\" \"comment_to\" " +
		"subtractfeefromamount )\n" +
		"\nSend an amount to a given address.\n" +
		"\nArguments:\n" +
		"1. \"address\"            (string, required) The bitcoin address " +
		"to send to.\n" +
		"2. \"amount\"             (numeric or string, required) The " +
		"amount in " +
		util.CurrencyUnit +
		" to send. eg 0.1\n" +
		"3. \"comment\"            (string, optional) A comment used to " +
		"store what the transaction is for. \n" +
		"                             This is not part of the transaction, " +
		"just kept in your wallet.\n" +
		"4. \"comment_to\"         (string, optional) A comment to store " +
		"the name of the person or organization \n" +
		"                             to which you're sending the " +
		"transaction. This is not part of the \n" +
		"                             transaction, just kept in your " +
		"wallet.\n" +
		"5. subtractfeefromamount  (boolean, optional, default=false) The " +
		"fee will be deducted from the amount being sent.\n" +
		"                             The recipient will receive less " +
		"bitcoins than you enter in the amount field.\n" +
		"\nResult:\n" +
		"\"txid\"                  (string) The transaction id.\n" +
		"\nExamples:\n" +
		HelpExampleCli("sendtoaddress", "\"1M72Sfpbz1BPpXFHz9m3CdqATR44Jvaydd\"", "0.1") +
		HelpExampleCli("sendtoaddress", "\"1M72Sfpbz1BPpXFHz9m3CdqATR44Jvaydd\"", "0.1",
			"\"donation\"", "\"seans outpost\"") +
		HelpExampleCli("sendtoaddress", "\"1M72Sfpbz1BPpXFHz9m3CdqATR44Jvaydd\"", "0.1",
			"\"\"", "\"\"", "true") +
		HelpExampleRPC("sendtoaddress", "\"1M72Sfpbz1BPpXFHz9m3CdqATR44Jvaydd\"", "0.1",
			"\"donation\"", "\"seans outpost\"")

	getbalanceDesc = "getbalance ( \"account\" minconf include_watchonly )\n" +
		"\nIf account is not specified, returns the server's total " +
		"available balance.\n" +
		"If account is specified (DEPRECATED), returns the balance in the " +
		"account.\n" +
		"Note that the account \"\" is not the same as leaving the " +
		"parameter out.\n" +
		"The server total may be different to the balance in the default " +
		"\"\" account.\n" +
		"\nArguments:\n" +
		"1. \"account\"         (string, optional) DEPRECATED. The account " +
		"string may be given as a\n" +
		"                     specific account name to find the balance " +
		"associated with wallet keys in\n" +
		"                     a named account, or as the empty string " +
		"(\"\") to find the balance\n" +
		"                     associated with wallet keys not in any named " +
		"account, or as \"*\" to find\n" +
		"                     the balance associated with all wallet keys " +
		"regardless of account.\n" +
		"                     When this option is specified, it calculates " +
		"the balance in a different\n" +
		"                     way than when it is not specified, and which " +
		"can count spends twice when\n" +
		"                     there are conflicting pending transactions " +
		"temporarily resulting in low\n" +
		"                     or even negative balances.\n" +
		"                     In general, account balance calculation is " +
		"not considered reliable and\n" +
		"                     has resulted in confusing outcomes, so it is " +
		"recommended to avoid passing\n" +
		"                     this argument.\n" +
		"2. minconf           (numeric, optional, default=1) Only include " +
		"transactions confirmed at least this many times.\n" +
		"3. include_watchonly (bool, optional, default=false) Also include " +
		"balance in watch-only addresses (see 'importaddress')\n" +
		"\nResult:\n" +
		"amount              (numeric) The total amount in " +
		util.CurrencyUnit + " received for this account.\n" +
		"\nExamples:\n" +
		"\nThe total amount in the wallet\n" +
		HelpExampleCli("getbalance") +
		"\nThe total amount in the wallet at least 5 blocks confirmed\n" +
		HelpExampleCli("getbalance", "\"*\"", "6") +
		"\nAs a json rpc call\n" +
		HelpExampleRPC("getbalance", "\"*\"", "6")

	gettransactionDesc = "gettransaction \"txid\" ( include_watchonly )\n" +
		"\nGet detailed information about in-wallet transaction <txid>\n" +
		"\nArguments:\n" +
		"1. \"txid\"                  (string, required) The transaction " +
		"id\n" +
		"2. \"include_watchonly\"     (bool, optional, default=false) " +
		"Whether to include watch-only addresses in balance calculation " +
		"and details[]\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"amount\" : x.xxx,        (numeric) The transaction amount " +
		"in " +
		util.CurrencyUnit +
		"\n" +
		"  \"fee\": x.xxx,            (numeric) The amount of the fee in " +
		util.CurrencyUnit +
		". This is negative and only available for the \n" +
		"                              'send' category of transactions.\n" +
		"  \"confirmations\" : n,     (numeric) The number of " +
		"confirmations\n" +
		"  \"blockhash\" : \"hash\",  (string) The block hash\n" +
		"  \"blockindex\" : xx,       (numeric) The index of the " +
		"transaction in the block that includes it\n" +
		"  \"blocktime\" : ttt,       (numeric) The time in seconds since " +
		"epoch (1 Jan 1970 GMT)\n" +
		"  \"txid\" : \"transactionid\",   (string) The transaction id.\n" +
		"  \"time\" : ttt,            (numeric) The transaction time in " +
		"seconds since epoch (1 Jan 1970 GMT)\n" +
		"  \"timereceived\" : ttt,    (numeric) The time received in " +
		"seconds since epoch (1 Jan 1970 GMT)\n" +
		"  \"bip125-replaceable\": \"yes|no|unknown\",  (string) Whether " +
		"this transaction could be replaced due to BIP125 " +
		"(replace-by-fee);\n" +
		"                                                   may be unknown " +
		"for unconfirmed transactions not in the mempool\n" +
		"  \"details\" : [\n" +
		"    {\n" +
		"      \"account\" : \"accountname\",      (string) DEPRECATED. " +
		"The account name involved in the transaction, can be \"\" for the " +
		"default account.\n" +
		"      \"address\" : \"address\",          (string) The bitcoin " +
		"address involved in the transaction\n" +
		"      \"category\" : \"send|receive\",    (string) The category, " +
		"either 'send' or 'receive'\n" +
		"      \"amount\" : x.xxx,                 (numeric) The amount " +
		"in " +
		util.CurrencyUnit + "\n" +
		"      \"label\" : \"label\",              " +
		"(string) A comment for the address/transaction, " +
		"if any\n" +
		"      \"vout\" : n,                       " +
		"(numeric) the vout value\n" +
		"      \"fee\": x.xxx,                     " +
		"(numeric) The amount of the fee in " +
		util.CurrencyUnit +
		". This is negative and only available for the \n" +
		"                                           'send' category of " +
		"transactions.\n" +
		"      \"abandoned\": xxx                  (bool) 'true' if the " +
		"transaction has been abandoned (inputs are respendable). Only " +
		"available for the \n" +
		"                                           'send' category of " +
		"transactions.\n" +
		"    }\n" +
		"    ,...\n" +
		"  ],\n" +
		"  \"hex\" : \"data\"         (string) Raw data for transaction\n" +
		"}\n" +
		"\nExamples:\n" +
		HelpExampleCli("gettransaction", "\"1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d\"") +
		HelpExampleCli("gettransaction", "\"1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d\"", "true") +
		HelpExampleRPC("gettransaction", "\"1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d\"")

	sendmanyDesc = "sendmany \"fromaccount\" {\"address\":amount,...} ( minconf \"comment\" [\"address\",...] )\n" +
		"\nSend multiple times. Amounts are double-precision floating point numbers." +
		"\nArguments:\n" +
		"1. \"fromaccount\"         (string, required) DEPRECATED. The account to send the funds from. " +
		"Should be \"\" for the default account\n" +
		"2. \"amounts\"             (string, required) A json object with addresses and amounts\n" +
		"    {\n" +
		"      \"address\":amount   (numeric or string) The bitcoin address is the key, " +
		"the numeric amount (can be string) in " + util.CurrencyUnit + " is the value\n" +
		"      ,...\n" +
		"    }\n" +
		"3. minconf                 (numeric, optional, " +
		"default=1) Only use the balance confirmed at least this many times.\n" +
		"4. \"comment\"             (string, optional) A comment\n" +
		"5. subtractfeefrom         (array, optional) A json array with addresses.\n" +
		"                           The fee will be equally deducted from " +
		"the amount of each selected address.\n" +
		"                           Those recipients will receive less " +
		"bitcoins than you enter in their corresponding amount field.\n" +
		"                           If no addresses are specified here, the sender pays the fee.\n" +
		"    [\n" +
		"      \"address\"          (string) Subtract fee from this address\n" +
		"      ,...\n" +
		"    ]\n" +
		"\nResult:\n" +
		"\"txid\"                   (string) The transaction id for the send. " +
		"Only 1 transaction is created regardless of \n" +
		"                                    the number of addresses.\n" +
		"\nExamples:\n" +
		"\nSend two amounts to two different addresses:\n" +
		HelpExampleCli("sendmany", "\"\"",
			"\"{\\\"1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX\\\":0.01,\\\"1353tsE8YMTA4EuV7dgUXGjNFf9KpVvKHz\\\":0.02}\"") +
		"\nSend two amounts to two different addresses setting the " +
		"confirmation and comment:\n" +
		HelpExampleCli("sendmany", "\"\"",
			"\"{\\\"1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX\\\":0.01,\\\"1353tsE8YMTA4EuV7dgUXGjNFf9KpVvKHz\\\":0.02}\"",
			"6", "\"testing\"") +
		"\nSend two amounts to two different addresses, subtract fee from " +
		"amount:\n" +
		HelpExampleCli("sendmany", "\"\"",
			"\"{\\\"1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX\\\":0.01,\\\"1353tsE8YMTA4EuV7dgUXGjNFf9KpVvKHz\\\":0.02}\"",
			"1", "\"\"",
			"\"[\\\"1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX\\\",\\\"1353tsE8YMTA4EuV7dgUXGjNFf9KpVvKHz\\\"]\"") +
		"\nAs a json rpc call\n" +
		HelpExampleRPC("sendmany", "\"\"",
			"\"{\\\"1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX\\\":0.01,\\\"1353tsE8YMTA4EuV7dgUXGjNFf9KpVvKHz\\\":0.02}\"",
			"6", "\"testing\"")

	fundrawtransactionDesc = "fundrawtransaction \"hexstring\" ( options )\n" +
		"\nAdd inputs to a transaction until it has enough in value to " +
		"meet its out value.\n" +
		"This will not modify existing inputs, and will add at most one " +
		"change output to the outputs.\n" +
		"No existing outputs will be modified unless " +
		"\"subtractFeeFromOutputs\" is specified.\n" +
		"Note that inputs which were signed may need to be resigned after " +
		"completion since in/outputs have been added.\n" +
		"The inputs added will not be signed, use signrawtransaction for " +
		"that.\n" +
		"Note that all existing inputs must have their previous output " +
		"transaction be in the wallet.\n" +
		"Note that all inputs selected must be of standard form and P2SH " +
		"scripts must be\n" +
		"in the wallet using importaddress or addmultisigaddress (to " +
		"calculate fees).\n" +
		"You can see whether this is the case by checking the \"solvable\" " +
		"field in the listunspent output.\n" +
		"Only pay-to-pubkey, multisig, and P2SH versions thereof are " +
		"currently supported for watch-only\n" +
		"\nArguments:\n" +
		"1. \"hexstring\"           (string, required) The hex string of " +
		"the raw transaction\n" +
		"2. options                 (object, optional)\n" +
		"   {\n" +
		"     \"changeAddress\"          (string, optional, default pool " +
		"address) The bitcoin address to receive the change\n" +
		"     \"changePosition\"         (numeric, optional, default " +
		"random) The index of the change output\n" +
		"     \"includeWatching\"        (boolean, optional, default " +
		"false) Also select inputs which are watch only\n" +
		"     \"lockUnspents\"           (boolean, optional, default " +
		"false) Lock selected unspent outputs\n" +
		"     \"reserveChangeKey\"       (boolean, optional, default true) " +
		"Reserves the change output key from the keypool\n" +
		"     \"feeRate\"                (numeric, optional, default not " +
		"set: makes wallet determine the fee) Set a specific feerate (" +
		util.CurrencyUnit +
		" per KB)\n" +
		"     \"subtractFeeFromOutputs\" (array, optional) A json array of " +
		"integers.\n" +
		"                              The fee will be equally deducted " +
		"from the amount of each specified output.\n" +
		"                              The outputs are specified by their " +
		"zero-based index, before any change output is added.\n" +
		"                              Those recipients will receive less " +
		"bitcoins than you enter in their corresponding amount field.\n" +
		"                              If no outputs are specified here, " +
		"the sender pays the fee.\n" +
		"                                  [vout_index,...]\n" +
		"   }\n" +
		"                         for backward compatibility: passing in a " +
		"true instead of an object will result in " +
		"{\"includeWatching\":true}\n" +
		"\nResult:\n" +
		"{\n" +
		"  \"hex\":       \"value\", (string)  The resulting raw " +
		"transaction (hex-encoded string)\n" +
		"  \"fee\":       n,         (numeric) Fee in " +
		util.CurrencyUnit + " the resulting transaction pays\n" +
		"  \"changepos\": n          (numeric) The " +
		"position of the added change output, or -1\n" +
		"}\n" +
		"\nExamples:\n" +
		"\nCreate a transaction with no inputs\n" +
		HelpExampleCli("createrawtransaction", "\"[]\"", "\"{\\\"myaddress\\\":0.01}\"") +
		"\nAdd sufficient unsigned inputs to meet the output value\n" +
		HelpExampleCli("fundrawtransaction", "\"rawtransactionhex\"") +
		"\nSign the transaction\n" +
		HelpExampleCli("signrawtransaction", "\"fundedtransactionhex\"") +
		"\nSend the transaction\n" +
		HelpExampleCli("sendrawtransaction", "\"signedtransactionhex\"")

	addmultisigaddressDesc = "addmultisigaddress nrequired [\"key\",...] ( \"account\" )\n" +
		"\nAdd a nrequired-to-sign multisignature address to the wallet.\n" +
		"Each key is a Bitcoin address or hex-encoded public key.\n" +
		"If 'account' is specified (DEPRECATED), assign address to that " +
		"account.\n" +
		"\nArguments:\n" +
		"1. nrequired        (numeric, required) The number of required " +
		"signatures out of the n keys or addresses.\n" +
		"2. \"keys\"         (string, required) A json array of bitcoin " +
		"addresses or hex-encoded public keys\n" +
		"     [\n" +
		"       \"address\"  (string) bitcoin address or hex-encoded " +
		"public key\n" +
		"       ...,\n" +
		"     ]\n" +
		"3. \"account\"      (string, optional) DEPRECATED. An account to " +
		"assign the addresses to.\n" +
		"\nResult:\n" +
		"\"address\"         (string) A bitcoin address associated with " +
		"the keys.\n" +
		"\nExamples:\n" +
		"\nAdd a multisig address from 2 addresses\n" +
		HelpExampleCli("addmultisigaddress", "2",
			"\"[\\\"16sSauSf5pF2UkUwvKGq4qjNRzBZYqgEL5\\\",\\\"171sgjn4YtPu27adkKGrdDwzRTxnRkBfKV\\\"]\"") +
		"\nAs json rpc call\n" +
		HelpExampleRPC("addmultisigaddress", "2",
			"\"[\\\"16sSauSf5pF2UkUwvKGq4qjNRzBZYqgEL5\\\",\\\"171sgjn4YtPu27adkKGrdDwzRTxnRkBfKV\\\"]\"")
)
