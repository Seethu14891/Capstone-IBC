// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DogPetMarket {
    address payable public owner;
    uint256 public productCount = 0;
    uint256 public dogPetCount = 0;

    struct Product {
        uint256 id;
        string name;
        uint256 price;
        address payable owner;
        bool purchased;
    }

    struct DogPet {
        uint256 id;
        string name;
        uint256 price;
        address payable owner;
        bool purchased;
    }

    mapping(uint256 => Product) public products;
    mapping(uint256 => DogPet) public dogPets;

    event ProductCreated(
        uint256 id,
        string name,
        uint256 price,
        address payable owner,
        bool purchased
    );

    event DogPetCreated(
        uint256 id,
        string name,
        uint256 price,
        address payable owner,
        bool purchased
    );

    event ProductPurchased(
        uint256 id,
        string name,
        uint256 price,
        address payable owner,
        bool purchased
    );

    event DogPetPurchased(
        uint256 id,
        string name,
        uint256 price,
        address payable owner,
        bool purchased
    );

    constructor() {
        owner = payable(msg.sender);
    }

    function createProduct(string memory _name, uint256 _price) public {
        require(bytes(_name).length > 0, "Name cannot be empty");
        require(_price > 0, "Price must be greater than 0");

        productCount++;
        products[productCount] = Product(
            productCount,
            _name,
            _price,
            payable(msg.sender),
            false
        );

        emit ProductCreated(productCount, _name, _price, payable(msg.sender), false);
    }

    function createDogPet(string memory _name, uint256 _price) public {
        require(bytes(_name).length > 0, "Name cannot be empty");
        require(_price > 0, "Price must be greater than 0");

        dogPetCount++;
        dogPets[dogPetCount] = DogPet(
            dogPetCount,
            _name,
            _price,
            payable(msg.sender),
            false
        );

        emit DogPetCreated(dogPetCount, _name, _price, payable(msg.sender), false);
    }

    function purchaseProduct(uint256 _id) public payable {
        Product memory _product = products[_id];
        require(_product.id > 0 && _product.id <= productCount, "Invalid product id");
        require(!_product.purchased, "Product has already been purchased");
        require(msg.value >= _product.price, "Insufficient funds");

        _product.owner.transfer(msg.value);
        _product.owner = payable(msg.sender);
        _product.purchased = true;
        products[_id] = _product;

        emit ProductPurchased(_product.id, _product.name, _product.price, payable(msg.sender), true);
    }

    function purchaseDogPet(uint256 _id) public payable {
        DogPet memory _dogPet = dogPets[_id];
        require(_dogPet.id > 0 && _dogPet.id <= dogPetCount, "Invalid dog pet id");
        require(!_dogPet.purchased, "Dog pet has already been purchased");
        require(msg.value >= _dogPet.price, "Insufficient funds");

        _dogPet.owner.transfer(msg.value);
        _dogPet.owner = payable(msg.sender);
        _dogPet.purchased = true;
        dogPets[_id] = _dogPet;

        emit DogPetPurchased(_dogPet.id, _dogPet.name, _dogPet.price, payable(msg.sender), true);
    }
}
