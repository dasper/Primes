# Primes Benchmark

Simple benchmark used to compare languages and to familiarize poeple to other languages. Fork based off of the code used in this youtube video:
https://www.youtube.com/watch?v=D3h62rgewZM

## Criteria
- MUST be native to the language being used
    - SHOULD NOT use additional modules, non-native JITs, or Exploits as these defeat the intention of using the code as a way to famalize people to a new language.
- MUST be of quality that would be deployed to a production environment
- MUST utilize [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes) for comparitive measure
- Sieve storage mechinism MUST be cleared/reset per itteration
    - This helps show examples of memory management as well as shows the performance gains/losses of memory management/garbage collection
- Benchmark timer MUST remain active through entire duration
