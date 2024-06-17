//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import { ERC721 } from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract DummyArtBlocks is ERC721, Ownable {

    uint256 public curTokenId = 13000000;
    constructor() ERC721("Dummy ArtBlocks", "DART")  {}
    function mint(address owner) public {
        _safeMint(owner,curTokenId);
        curTokenId++;
    }
	function multimint(address owner,uint256 num) external {
		for (uint256 i = 0; i<num; i++) {
			mint(owner);
		}
	}
	function _baseURI() internal view override returns (string memory) {
		return "https://token.artblocks.io/";
	}
}
