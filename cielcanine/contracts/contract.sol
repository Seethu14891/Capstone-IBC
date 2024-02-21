// SPDX-License-Identifier: MIT 

pragma solidity ^0.8.0;




contract sample{
    event MessageLogged(string message);

    function logHelloWorld() public {
        emit MessageLogged("Hello World!");
    }
}