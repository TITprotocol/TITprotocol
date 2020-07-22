// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

var MainnetBootnodes = []string{
	"enode://a65b59893533c12779669fbac0c9b8c362afc12c79b029655d1d74b343bb58f4736a8c140367cf9262fe7e9f46a7db4a595a2812e5cde115acb84283c4bfc0b3@13.113.138.251:30303",
	"enode://74be4fdc203fc93116ec1f914d02b7935f9774696a2550a896e638d7d79ba2ffe74778e3e2eabb2394355293e2df54fc7ec256143c2f25526d0bd577f27280fd@52.198.25.200:30303",
	"enode://c864e9c5cd7035107565ae5ac74bb408192183569422163a98c5b28f34e8ea1a7a4211e03b1189d62a548c3a40563043eccd53fcf0bdbc7d37fac5c948d0bfda@18.180.127.112:30303",
}

var MainnetInitIds = []string{
	"2eb904348b2dbebd",
	"f45eb7532b909e75",
	"a113b2d4641a8200",
	"54093c2d160a33d7",
	"cf99df143cd12bb2",
	"8a29ab8445921e58",
	"2194ec5b50be4ac1",
	"5d1a2bc8be2a9b54",
	"f6d1cdc6c5511fa8",
	"2aafeadb97f17f41",
	"69c940a937610926",
	"dde515a022829b23",
	"8fa26e333d58a261",
	"b85c4d70118a7ff9",
	"e5fdca6084db72b6",
	"35cdea3ad9267f2c",
	"711db9db4569c7d4",
	"03e3e08a2f5feec9",
	"987e1f2bfce81a8d",
	"5a5ba3f57d9d7e80",
	"1f749971e3a11d41",
	"52ff1119f748ef9d",
	"17c20301b93c9d84",
	"cb531e58ad81e094",
	"f59da46f7f30357b",
	"6e2f68f520040628",
	"879eec2d083bf692",
	"92a28f69c521117f",
	"0bfe75449d6ee026",
	"45b476a588fce8a8",
	"fcd3db9d4ff3a5ed",
	"63cde7616c1ab53a",
	"65ba3e3600b41e64",
	"9a702f5cbfe24456",
	"ecc7bbfaca0daabb",
	"f09c6d2d3c461005",
	"a6ab67a5744f744a",
	"0d99d74cf15ef0f7",
	"ccdfd3e1e1b7ed0c",
	"850558bbb0ba3749",
	"16474e9ae22a36e0",
	"525a81c948ff5603",
	"4bc53cdebc207b48",
	"03ad95102d3dec50",
	"54dc0a2b2d305bbb",
	"e4dff7f3b39faa6b",
	"e0b827d4e3048baa",
	"38cdc6a1bb85d527",
	"f00fdddd02204342",
	"227bba779040ae5a",
	"13c12c785b3b1ee1",
	"15a7f4e64baaba6e",
	"b0b03c98e42e8059",
	"879506759a55d20f",
	"56835c49d493c46b",
	"ab932a34ac8bb685",
	"99cbc55ea65adff6",
	"c9dddf83645ad68a",
	"19fd11aabfdc9cb9",
	"7809fb32bb86ed87",
}

var TestnetBootnodes = []string{
	"enode://59ca967b2c9c1442e81026f5ffc2b24f4b3787512194a41e4ab14dfac97e75b700988cac80f973641d40cd65f775f41955b93d2e843ebb03555b16dd9bf983d4@127.0.0.1:9646",
}

var TestnetMasternodes = []string{
	"enode://2eb904348b2dbebd1eae813067cea49ab127cab3efafb38bdefbb641b5b36938ece0d76b35b516634eae7de0c3034dbac2bd1f39a325160711a6eebdcbc41286",
	"enode://f45eb7532b909e75d361d41276a3074e8e83636c8f1f169db90c2cc2f4f525ddf3ade2366c8c8369e7763fba80fab580023843ee8a5710328bbe4e3a99c453d2",
	"enode://a113b2d4641a82007ed452751bab2fcc78ee16663f34995618ccea96d8e396a704c28f471e52c86f1857db04c07ca6be2ae6f8b1015ebea4b16208429fbd950d",
	"enode://54093c2d160a33d746a7c963fa95b9f67e836934e4ac932a4af1ecf6889265316863d54d78c269e271ce7ea02d462f4d1149cd1de6aec87eb3fc6c424778bf4d",
	"enode://cf99df143cd12bb2e8f98acf13c571e7272e2e049f582b8c5f08a526c8c79d083d297e6aac804aa966a87857e3141fd230d6eae0e24ff361a7d5871da98a16de",
	"enode://8a29ab8445921e5823901a9f5663da85cf0a9fb0aeb50f9677e35300b58583a9315ee3292972c035df041766721884858a27dbebb947f0d720f00752b853e7fc",
	"enode://2194ec5b50be4ac19b60918dbb682dff362a1b9d6348b161afed156ce162c94248d254e58ff47f2874c01f47561d4bd0827981efd9f0b11cdde3d2d54daddf54",
	"enode://5d1a2bc8be2a9b544e592fadbe2ac30c20e8f8b6996c4c0e698ba2b4d3a09427afa674dc523a87b0afed6f24abdc72412bd88ede29207a117ebbdaf95cce4530",
	"enode://f6d1cdc6c5511fa8b877bd60223467bff0b2caf476565d7c82d3867a1b10633d1b09dc939d277531e0ad786056881244ec37cdbcb1ca53cce457e257949605dc",
	"enode://2aafeadb97f17f413d464d18235f56117060a56d42ff9a96414153e19887355405ebd3eef68a78ed8c672d31212c317b1aafa4d7e0cc020c6c7675e496f4f54d",
	"enode://69c940a937610926e82c44b6c5d8a6c8079201e090dc54afed2bd5ed91790d289130dbb820984f2f170c119219b99391360f499f08b7535b32c02dc1507e3574",
	"enode://dde515a022829b233cd6c770e94309df17863d4d305430f4d640408460ccf408024052529ca4e22e6e6ada42690a768da0c9c575062012ac40a9fda0f3067e11",
	"enode://8fa26e333d58a2611e4302db4befb5b6aa68475eb0d78a92572fca21fc95d8564373076b2248284db4b670b21e1f60855faa7df78934c017ca9c423f606fcfce",
	"enode://b85c4d70118a7ff957cfa46b2874bab40b543d3161d5526f8ee489a414de719bdcdb86f82c891af9e6609250c70620bd355664ac8d2405bc89c89eaac07a2c17",
	"enode://e5fdca6084db72b652b5ae4dfb99df8dd1de476e70df6e9236e4f6e5d8c73ef1d90380d5cca8f7e2c964eb8625eaac4c03477db693412937549cadfd34a05244",
	"enode://35cdea3ad9267f2c969f087a8310ac5209657c006fb99c5600f5050d9b9805ce9b929f2ae983442867d5e2a53d41962963e1d262875c9fde0ff87f99150ed9bd",
	"enode://711db9db4569c7d4085eab514d92422a9e22958c5f2c5feb8c092a2b859d14c5785c8652e4d96aac0af8ee29e2d6d4fe638b4e35deb0d3ae3dec939ded92ed7f",
	"enode://03e3e08a2f5feec99b022dcdda5c713e28a2cc0a79057c43625d8aa9c286ae18836362d793ae80e791420454d8df576fd1d7c58e85637026743bb85f9752c35e",
	"enode://987e1f2bfce81a8d0016563060e268a9e01548621f0f339e152081e5a807d65e5a4f75e68c714a7f87e74f8282a64a72a75725b1b1ee0ce7584b7bb7a3473ea8",
	"enode://5a5ba3f57d9d7e8011406d6f75277aee310e3852371e64676a775d6cb30865838ae78241fac5b2e24dab76ea846f5f8a68f6f6d1fc4496f48c7469677cdda76e",
	"enode://1f749971e3a11d414b03677a81204373af25bb1769760355f733ab19f0f8590385cbcab15e8277f0e82436b16c777db78ddf4e93c8031385491c5cb93a053acb",
	"enode://52ff1119f748ef9d93c268df1fc0a7ab2bfde438aa6b20f12787df2d76ca17635773d69d8ba1c797291eebd9da69aacd77ae43c0e66ae65e25eae37bbc856c01",
	"enode://17c20301b93c9d84eed82082ee7394ece7cf062e2d46901fc48ee30832102189c05843190615ed3bea613e5491c2a2021b490ed0c7ab43d52ed4ae238402064c",
	"enode://cb531e58ad81e094e8be5135efe3b69fe6f15769e8a25eb6662bee6c5793e12f863ef64cdade32a79e40227c31486e360a29062c6912564a605c5f821e0b5d7a",
	"enode://f59da46f7f30357b3f131a7685ac90f8f57b6d7afc6021a6d332294f829a1be10f6d201d9f68faa767e4365051dc358b03a81577f4dcb37bdade87dff270e988",
	"enode://6e2f68f520040628fb4b245b5ed174f8c3d8a4f2b8b8c9a20e1db002f94e8fd3c22434baa6f786bda2e2ecad340c42c3586eec0ba9aac317ca23e4bcec2d4590",
	"enode://879eec2d083bf6927a29d04d7414ae243efdc087a74ab634bcf173e27cea8f9126f11f7773afe331163101e7d4ebe8c59df24902855945d43825f021b9ff0cea",
	"enode://92a28f69c521117fea90a175346fc3750dd218954870e9839e7d7a435073edda93472b16fdd0ca7d1b79663ccc647cddebce3ec8d0678564161d093cc8ee6104",
	"enode://0bfe75449d6ee026f118a0dcc657402714f20eb18e5bdda7e5e8186dd603cf05a91dc095aa7960666179382b4becaf40535123d9e03ee2c29f23e215c5b3c97a",
	"enode://45b476a588fce8a8a933f6f621e22679589451966e8b7b871182680de9b414a89ebfe1b846a4ccbda5d9a79319a9a462be343dfff1564501df747442df86fe85",
	"enode://fcd3db9d4ff3a5eda8b5a08ee0dfc9ea483111a51e893f39889a697bdff18e353ac20f711328155e24811d6ef23f2867d30e676c106f883190003a67b9c08ab0",
	"enode://63cde7616c1ab53adadb0641a4d3a4509d47881597f59ea472767cfdcd298f8d10c7d289303260226a613d2577349a490366d5f8eb689e72885b0c99c6c6c165",
	"enode://65ba3e3600b41e64b089004aeeb879bd55c7253567abd1db4ab8c9d47fda72cde24cf6fa1343cffbd0de46c7a8dd96206fb6ba8bb22eaa2bbd310ab935dee14d",
	"enode://9a702f5cbfe244561394c720fce199031ac4da7c51a646de1f9eeb95f6e0104075845b8b713dd1ee07c973c7f0abba4180b75c5df0dc294aab7331398947261b",
	"enode://ecc7bbfaca0daabb90c54f7dde68add2b70bcdca75485e0e1b762f6126ffb3b746aff8df2e538664110a0f9836b822a939aba55f58fe619e0d8f185085bbff9e",
	"enode://f09c6d2d3c4610050faf4a0e776fb8ca4cfaca3006d9731c61357386e11f05b6f6f204f28de697289d9669a0a827dde6779325b5290ea7ad89b450d25f8f997c",
	"enode://a6ab67a5744f744af2a2b5cc6233f328c7ec3c3b81c6cd6b233149c1f9388085b876d375af5db4a10ef940ccdb15f0747d52705752a4446129d03ed9e961ecdf",
	"enode://0d99d74cf15ef0f7cdef0bb722f0155f74482d44a7fd2a787b68cb62f4f600887d7caef4bcb3c7052f68beffdb3440d28d47cdb107850c54977c4f6736ae4e83",
	"enode://ccdfd3e1e1b7ed0c107bd555f334a806adfe55f4fedea43d6c027c463196563671efc175fd6ddf4a85da21542e8f3e104c510a2a906c39fc125637b301860b85",
	"enode://850558bbb0ba37490ebef4987b77906cdd23fc82345c844a2915c47e0afa2db3417f2209908200a5185b83909c28154b182297a80887f08d338d8d23a32cc5a4",
	"enode://16474e9ae22a36e00e55e9ba4d45ce0fa392fedf76fc164d6446b9fbef59e98c2a14b79ab703f23da25287526c5c1f4ce5e4f917973d590e3e9c6ac011bfc2a3",
	"enode://525a81c948ff560309d0257d7cf1205efb4fe664381b808c6df449c76432e555bd17dca1c4b12c250a56bf7a25a3de9d7526a007ec379a764d7add24ca2fc3d7",
	"enode://4bc53cdebc207b4853515f99e4e810ff4396131399b5fd0d577b4e7aebedd20adcb7be9cfdeb1cc589e46d325e4a2b2a04434df4e9e69cd05c3396a7ad53b57d",
	"enode://03ad95102d3dec5017362756c62339f776acffb66a4b889a2651e1b7d59ab3228105d7e0ed3f4a7f3d078418122df0de69833746088e7401ca8862dea9bc31b1",
	"enode://54dc0a2b2d305bbb88c10b691b9fee4b272cc98090ac2451095048974cc5ee95196c60c4652ddb253ac27de7d7bf86c2e6d6311f63354b303ec579c143a2bc67",
	"enode://e4dff7f3b39faa6b8ab8a8fb276d75d39e157e00813d401b9631b543a308bbf0c510cae771330aa6b455d46c96ecf916bfbc0f24ae48006ba51865f2e28f1f65",
	"enode://e0b827d4e3048baa5821c0d748ffe1fcd10357d5ce6044c6b1e2a32e93810dcca421e257022fd23bad14dfe0b7007bddb5801f8f2e86a1aa6c6a430ed544330b",
	"enode://38cdc6a1bb85d527239abb547e5cc2e8b03a55be7a016763ffc879df4835ff03ea8c385c8978ce9b1c752915470d3b96d77d6c36ed5aab42efc589d6a72b76df",
	"enode://f00fdddd022043428ae25074bfe8c0d8b8ce4a3fd7607e81fe049d812a02fd29113a32b4fe989202990a8bb7dc6c7efc51c4004c34c8174a469ec734db36e97e",
	"enode://227bba779040ae5a82420bb64df38f7f69a5a2185d603c08916a186fa484b628be9c5fa69935cdc9c112c1ce01b9328714f2576e0ed7fe00bf621fbfb68826ba",
	"enode://13c12c785b3b1ee12e9c6ec66642383ecffb17c122518c006c1b0a6c1ee7bb254ceda56e66d7a21b4b5c0684735587ee79af6964c43e176b9852763522fcf091",
	"enode://15a7f4e64baaba6ea7edd814d86d4496efbad4b49a828191a5e8138ae54dd4baa39d74f473f8fcb0e436e072c0215e548dba27992db0b0c30d6ee454aaa10661",
	"enode://b0b03c98e42e8059fc568fdb85c284d29eb8d6a95ebc1b3c6eaad8bbabd64a647907fba508549a2ba4db0a599d1fc524a57602f2989c78a273f81e2a8579c2f0",
	"enode://879506759a55d20fb09c464b4c3ce121b57a2fbedbef18913d41d4d8296f57399b2a63cc5d0a39bb7d284004b0f01431f04f2821ac51e9eeb63778e523eeac96",
	"enode://56835c49d493c46b26b4b1286fce4c908b920fce9bb0d9d109aa0e383968abec05bee082b273d8750eb399a48808bd66370186b97e4270f7bbdd922bc9017dab",
	"enode://ab932a34ac8bb685056e186fa02096eda1f9214c97274cba40ec8b370c990f2b497cd2d7b1ac26b1d2715ebfbdc72c0ce5779cf17fece7d51ea80abf5bc7a90c",
	"enode://99cbc55ea65adff61d6d997c271d1f6aa95a0a2ec37977726bcb383427990174daec1b2dac02e03b55a9b58b67b6835fd41ad1030a5234cc8a22b134646bff32",
	"enode://c9dddf83645ad68a9d314b205a19d57e68ffe97cc1dad2eadb2184cc3aebc6b8a5f575bd16e04920f462f1e75d031e1dada2341f725ff73fca8423791c0f25f5",
	"enode://19fd11aabfdc9cb9647791bb4ff509ec270b4b23d6c22c401daf5f65e2f6668638f6c70fafd3396753b22694054644478f8f1a14b844a2d02a5afaf9d2b969a9",
	"enode://7809fb32bb86ed87d83cab38048d3d955c72a2713a3ac2e183b603e45a40aa0e4ba6b8c6786b2f75a655e3a661e65f2a8d16158c245d05b6f19f53536792299e",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
