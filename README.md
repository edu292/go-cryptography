# Cryptography

## What is Cryptography?

To start off being a little pedantic, the meaning we normally attribute to cryptography is not entirely correct. In the dictionary sense, cryptography is the practice and study of creating secure communication techniques, but it doesn't refer to simply using or interacting with them. 

The term that encompasses the whole field is **cryptology**, the science of secret messages. Under cryptology, there is also another important field called **cryptanalysis**, which is the inverse of cryptography. It consists of the science of breaking and exploiting failures in cryptography systems. While this pedantic note is warranted, I will be using the common-sense meaning of "cryptography" in the rest of this readme's explanations for familiarity's sake.

### Encryption vs. Encoding
There are other misconceptions around cryptography: just because you can't make sense of something doesn't mean it is encrypted. Often, people see encoded text and assume it is secure. But for a system to actually be considered cryptography, it must not be easily reversible and must reveal as little as possible about the original information. Attributes that mere encoding does not have.

To be clear, everything other than straight binary relies on some type of **encoding**. Since computers only natively deal with 1s and 0s, representing anything else requires establishing a shared convention, pattern, or protocol. This is what an encoding is: a standardized way to attribute meaning to a specific value or a group of values in a specific pattern or position. In the human sense, we would be working with letters and attributing meaning to words and sentences to form a language. Because we are dealing with computers, that same concept is applied to ones and zeros, bits, and bytes.

We often stack multiple encodings on top of each other to represent complex data. If you look at that data without the proper program to translate it back to its human representation, it might look completely random and cryptic. However, once you know which "language" or format you are looking at and find the right program, any information there is as easy to read as anything out on the public internet.

Keeping this concept of encoding in mind is important for more than just preventing you from exposing sensitive data under a false sense of security. It serves two other key purposes for understanding and using cryptography:
1. **Mathematical Translation:** Cryptographic algorithms operate on the underlying binary representation of a message, sometimes treating it like a giant number to perform complex math, and encoding is what makes that translation possible. 
2. **Visual Representation:** The final result of a cryptographic algorithm (the **ciphertext**) is usually encoded again so it can be represented in a more concise or visual form. As you now know, this final layer of encoding adds zero security; it is strictly a matter of convenience.

## Types of Cryptography

Knowing what cryptography means is only the first part of the problem. Although translating text to and from gibberish might look one-dimensional, the sheer number of secure communication problems we face requires more than just one approach. There are several types of cryptography, and each has its specific usage, advantages, and trade-offs. The most important distinctions are:

* **Symmetric:** Methods of encryption that involve a single key. These tend to be the fastest and simplest to set up, especially when you are the one encrypting and using the data and don't need to share it with anyone else. It gets more complicated when you *do* need to share it, because once that key is transmitted through an insecure channel (like the internet), its security is gone. You have to find a way to safely share the key and trust everything it touches on its way there.
* **Asymmetric:** Encryption that uses a pair of keys (public and private). The public key can be exposed, and anyone with access to it can encrypt data, but *only* the private key is able to decrypt it. As long as the private key is kept safe, the system remains secure. Normally, this is used to sign data or to initially establish a secure channel. Due to its performance overhead and specific algorithmic constraints, it is not normally the main way bulk messages are encrypted.
* **Hashing:** Not actually an encryption method. The rule that the output must reveal as little as possible about the original information is still kept and is even more important for these algorithms. But, they are irreversible, meaning there is no key or method that can retrieve the plaintext back from a hash. Hash algorithms generate artifacts that could be described as a mathematical summary or fingerprint of a text or file. We can use them to attest that we have the right file and it hasn't been tampered with. They can also be combined with a secret key to prove that the data came from the expected source (MAC and HMAC), which is a different kind of signature we will cover later.

## History of Cryptography

While cryptography today is heavily associated with internet security, the need to protect intercepted messages is an ancient human challenge. 

### The Caesar Cipher and Brute Force
One of the earliest known attempts at secure communication is the **Caesar Cipher**. Attributed to Julius Caesar, this method consists of shifting all letters of the alphabet by a certain offset, rendering the message unintelligible unless you know the exact offset to shift the text back. 

Here we are introduced to one of the simplest methods of cracking encryption: **brute force**. Brute force involves trying all possible combinations until the correct one is found, ideally prioritizing the most likely options first. Because the English alphabet only has 26 letters, there are only 25 possible offsets before the shift wraps back around to the original text. With such a small pool of possibilities, this was a highly vulnerable cipher even in ancient Rome. Today, leveraging letter frequency analysis and modern processing power, a computer can crack a Caesar cipher in microseconds.

### The Enigma and Alan Turing
Another pivotal chapter in cryptography coincides with the very origin of computers, at least in the way we know them today. The **Enigma** was a German cryptographic machine relying on multiple electromechanical rotors and a plugboard to scramble messages with daily-rotating secrets. 

While Polish mathematicians had successfully cracked early versions of the Enigma before World War II, the German military's addition of more rotors and complex daily changes soon outpaced manual decryption. By the time a message could be cracked by hand, the intelligence or planned attacks would already be outdated. 

With the mission of finding a faster way to solve the Enigma, the British army recruited the top mathematical minds of the time, most notably **Alan Turing**. At the time, Turing was already a celebrated mathematician and logician who, under these fields, had published multiple relevant papers that would become the basis for all modern computer science and computer architecture. Building upon previous Polish concepts, Turing and his team determined that the only way to overcome Enigma's vast combination space would be creating a new kind of machine. This effort led to the creation of the **Bombe**, a machine that used logical deduction and known plaintext guesses to rapidly eliminate incorrect Enigma settings.

### The Impact on Modern Computing and Security
It is important to note that the Bombe did not fulfill all of Turing's ideas and concepts, and it would not be considered a true computer in the modern sense. It was a hardwired, single-purpose electromechanical decryption tool. However, its ability to automate cryptanalysis at unprecedented speeds had a massive impact on the outcome of the war. More importantly for our context, the Bombe served as a monumental proof of concept for the future of technology. It demonstrated that complex logical deduction could be successfully mechanized and reliably executed at scale. 

From a cryptographic perspective, the cracking of the Enigma highlighted a catastrophic vulnerability: **predictability**. The Bombe exploited a mechanical flaw in the Enigma—it could never encrypt a letter as itself—combined with the fact that operators frequently used predictable, highly structured formats, such as repeating the same greetings or weather headers. These known similarities ("cribs") allowed codebreakers to pin down the secret keys. 

This taught modern cryptographers a vital lesson: an attacker should never be able to deduce the key just because they know part of the plaintext. Building upon this knowledge, modern algorithms introduce additional techniques to add randomness:
* **Initialization Vectors (IVs):** Used in encryption to ensure that encrypting the exact same plaintext multiple times will always generate a completely different and randomized ciphertext.
* **Salts:** Used alongside hashing algorithms to prevent attackers from precomputing predictable outputs (like rainbow tables for passwords). 

Both techniques ensure that predictability is stripped from the equation, rendering pattern recognition and known-plaintext attacks useless.


## Applications of Cryptography
As we transition from theory to practice, you will notice a heavy bias toward hashing in the following applications. To be clear, it's impossible to quantify whether encryption or hashing is more "useful", they are simply two distinct tools required to build a secure system. But the reality is that applying standard encryption to a problem is usually straightforward: you lock the data and hide the key. Hashing is where the interesting engineering happens. Because we are intentionally destroying data to create a mathematical fingerprint, we have to rely on clever techniques, like stacking, salting, and stretching, to make that fingerprint useful.

### Hashmaps
They are in almost every language and can range from very important to the single most critical data structure in the codebase. You might know them as dictionaries (Python), objects (JavaScript) or maps(Go). 

Under the hood, they work by taking a key (like a string), running it through a hash function, and mathematically compacting that output to fit within the index bounds of an underlying array. This relies entirely on the strict determinism of hash functions, the same input key always yields the exact same hash, combined with the lightning-fast access speeds of array indices.

This is how languages instantly find variables in the global scope, how dynamic features like monkey patching and duck typing work, and it tends to be the go to optimization for a massive variety of algorithm problems. It is worth noting that while cryptographic hashes (like the ones used for MACs or Passwords) prioritize security and collision resistance, the internal hash functions used for hashmaps are explicitly designed to prioritize pure, unadulterated speed.

### Password Storage
Common-sense among web developers, and baked into most frameworks, is the concept of hashing passwords before storage. But, how does that work? if hashes are irreversible, how can you be sure that the passwords match when the user tries to log into their account? Wouldn't encryption be better for this? These are the questions to be solved in this section, and through them you will come to know a very important distiction between hash algorithms: Genereal Hashes and KDFs. 

The reason we use hashing is that we don't *ever* want these operations to be reversible. If someone were to gain access to the database and know the encryption secret, then all users' accounts would be exposed. And if access to the database has been gained, logically, we can also assume that the secret is known. While we can't reverse the stored password back to its original form, we can simply hash the input password before comparing them. Because hash algorithms are deterministic, meaning they always output the exact same result from the same input, this is how we can still know if the user's passwords match without ever needing to know the password itself. 

Despite this being a step in the right direction, two new issues arise. First, if hash algorithms are perfectly deterministic, then nothing stops the computation from being outsourced. Attackers can pre-compute the results for the most common passwords, store them, and re-use them on every database leak to instantly guess accounts. Second, general hash algorithms are, as the name implies, general. They tend to be designed with speed in mind, which is exactly what we *don't* want attackers to have when trying to brute-force user passwords. 

To solve these issues, we use salts and KDF algorithms. A salt is a random input passed along with the password into the hashing function (`Hash(password + salt)`), which changes the final output entirely, even for the exact same password and algorithm. Meanwhile, a Key Derivation Function (KDF) is a special type of hash function created specifically to consume more expensive computer resources or utilize operations that cannot be easily executed in parallel, purposely making mass hashing attempts incredibly difficult.

### Integrity and Authentication

#### Checksum
If a hash is a fingerprint of a file them it can be compared to the original file (by re-hashing it) to confirm whether there might've been an issue on the transport of the file, on a large download, for example. This technique is known as a **checksum**, and is very useful if the download and the hash come from different sources, or when you are absolutely sure you got the hash from a secure channel. 

However, as stated, it only works for accidental transport issues, not malicious manipulation. if an attacker is able to manipulate the file, they could simply generate a new hash for their corrupted file and swap that out, too.

#### MACs and HMACs
This is exactly the problem that MACs are designed to solve. They work by combining the checksum method with a secret key, appending it to the file before hashing it (`MAC = Hash(file + secret key)`). Because only you and the trusted source have this secret, we can rule out an attacker being able to manipulating the hash. Any attempt at doing so would be fail your hash comparison. 

This is still a specific use case, but there wouldn't be any issues if it weren't for **Length Extension Attacks**. This exploit targets a specific characteristic of popula hash functionsy like SHA-256. Because these functions process data sequentially in fixed-length blocks and rely on padding to make smaller blocks meet the size criteria, an attacker who knows original message's length can essentially "un-pad" it and make the function continue hashing any extra malicious payload they append. They can generate a valid hash and fool a standard MAC without even knowing the secret key. 

To solve this **HMAC** was created. It is basically a two step MAC (`HMAC = Hash(Hash(file + secret) + secret`), which mathematically seals the state and make length extension attacks impossible.

#### Digital Signatures
MACs and HMACs are highly effective, but they require both parties to securely share the exact same secret key. What if you want to prove a file is authentic to the entire public internet without giving everyone your secret? This is where **Digital Signatures** come in. 

Instead of a shared secret, it uses asymmetric encryption on top of the file hash, with a *Private Key* signing it(which is the same as encrypting). You then publish the file alongside this encrypted hash. Anyone can then use your widely known *Public Key* to decrypt and verify the hash. If the hashes match, it proves the file wasn't tampered with and that it absolutely had to come from you. 

#### CA Certificates
Digital Signatures solve the problem of public authenticity, but they introduce a critical "chicken and egg" problem: how do you know the Public Key you downloaded actually belongs to the person you think it does? If an attacker intercepts your connection, they could still hand you *their* public key while claiming it is Google's, allowing them to fake signatures. 

This is the vulnerability **CA Certificates** solve. A Certificate Authority (CA) is a globally trusted entity whose public keys are pre-installed and hardcoded directly into your operating system or web browser. When a server wants to prove its identity, it asks a CA to digitally sign its public key. Therefore, a CA Certificate is not a new cryptographic algorithm; it is simply a server's public key wrapped in a Digital Signature from a higher, pre-trusted authority. It acts as a cryptographic passport, chaining mathematical trust all the way up to the root.

#### JSON Web Tokens (JWTs) and Third-Party Auth
Every time a user wants to view their profile, send a message, or buy an item, the server needs to know who they are. We *could* just ask the user to send their password along with every single request, but that is obviously a terrible idea, it creates endless opportunities for those credentials to be intercepted or leaked. 

Instead, after first verifying the password, the server needs a way to say, "I verified this person, they are good to go." Traditionally, many websites handle this using session IDs stored in cookies, keeping track of who is logged in via a database on the server. However, another popular, stateless approach is to hand the user a cryptographically verifiable document known as a **JSON Web Token (JWT)**.

A JWT is just a standard JSON payload (usually containing a user's ID and role) with a signature appended to the end of it. When the user sends this token back on their next request, the server checks the signature to ensure the user didn't tamper with it (like changing their role to "admin"). 

When you are the one issuing these signatures and they are used only on your server, the simplest solution would be to use HMACs, since you wouldn't have to share the key used.

But what if you are implementing a feature like "Sign in with Google"? Your server doesn't share a secret key with Google, and you certainly don't have access to the user's Google password. When a user logs in via Google, Google authenticates them, generates a JWT, and signs it using their *Private Key*. Your server receives this token, fetches Google's widely available *Public Key*, and verifies the signature. If the math checks out, your server can securely log the user in. You are relying entirely on cryptographic proof of Google's identity, establishing secure authentication without ever touching the user's actual credentials.

