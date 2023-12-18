# CircuitBackend

The backend is used to verify that a user's passport photo corresponds to a real photo taken by a phone camera.

Also, it is used to store *Merkle trees* and generate inclusion proofs (*Merkle Branches & Roots*).

To run a server:

```console
make run
```

To run tests:

```console
make test
```

***Endpoints:***

- `/verifyPhoto` (**GET**)
    Receives photos. Uses EdDSA (BabyJubJub) + Poseidon Hash for signing the response in case of **successful** verification.
    Currently, the verification process is mocked and always returns **TRUE**.
  
   *Input format*:

   ```json
    {
        "PhotoReal" : {
            "ImageBytes" : [1]
        },
        "PhotoPassport" : {
            "ImageBytes" : [1]
        }
    }   
   ```

    *Output format*:

    ```json
    {
        "hashPassportPhoto": "18586133768512220936620570745912940619677854269274689475585506675881198879027",
        "hashRealPhoto": "18586133768512220936620570745912940619677854269274689475585506675881198879027",
        "photoHash": "1656567780223073447084190538735726676411629511500005572106155109614008362892",
        "signature": {
            "R8": {
                "X": 7672243323998928051244197237606963025417647751203971716427928900896151147300,
                "Y": 18832230028543216750284706231785595070962193420630427450228904590271894003990
            },
            "S": 952068376983978881253194807205172524444803223447058877048158962253044807588,
            "A": {
                "X": 4488698527135524487431810538112307395601028806365869678736730219632526477354,
                "Y": 13802359297867645163996080231069642651080118154952690766033165846082618832678
            }
        }
    }
    ```

- `/providerInclusionProof/:leafHash` (**GET**)
    The endpoint is used to receive a Merkle Branch for a provided Leaf. The Leaf itself is represented in the form of ***`Poseidon(A.X, A.Y)`***, where **A** is the public key.

    *Output format*:

    ```json
    {
        "branch": [
            12669517726012161444981800561236434280518113038714736880405414311681259007790
        ],
        "order": 2
    }
    ```

    *branch* array represents a Merkle Branch;
    *order* represents the order of leaves (left or right hashing with the branch value, `0` - *left*, `1` - *right*) if represented in the **binary form** from the **lowest** bits.
    For example 1000101 means *R(1), L(0), R(1), L(0), L(0), L(0)*. Currently, the first bit is always `1`. This is done to mark the tree's depth.

- `providerMerkleRoot` (**GET**)
    The endpoint is used to receive the *Merkle Root* of the providers' *Merkle Tree*.
     *Output format*

     ```json
    {
        "root": 12168667296552870753094296647991914965579574911497994731276783496022880154592
    }
     ```

- `providerList` (**GET**)
    The endpoint is used to receive the list of the providers identifiers (hashes of public keys)

    *Output format*

    ```json
        {
            "providers": [
                12669517726012161444981800561236434280518113038714736880405414311681259007790,
                12669517726012161444981800561236434280518113038714736880405414311681259007790
            ]
        }
    ```
