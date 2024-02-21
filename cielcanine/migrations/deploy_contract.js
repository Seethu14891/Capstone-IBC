
const DogPetMarketMigration = artifacts.require("DogPetMarket");

module.exports =  function(deployer) {
    deployer.deploy(DogPetMarketMigration);
}