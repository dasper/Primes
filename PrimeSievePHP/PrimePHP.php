<?php
// ---------------------------------------------------------------------------
// PrimePHP.php : Prime Sieve in PHP (7.4 and 8.0) implementation
// (as close as possible to the C++ implementation)
// ---------------------------------------------------------------------------

class PrimeSieve {
    private int $sieveSize = 0;
    private array $isPrime;
    const myDict = [
        10 => 4,        
        100 => 25,
        1000 => 168,
        10000 => 1229,
        100000 => 9592,
        1000000 => 78498,
        10000000 => 664579,
        100000000 => 5761455
    ];
 
    private function validateResults() : bool
    {
        return (self::myDict[$this->sieveSize]??null) === $this->countPrimes();
    }

    function __construct(int $n)
    {
        $this->sieveSize = $n;
        $this->createArray();
        $this->hydrateArray();
    }

    private function createArray()
    {
        $n = $this->sieveSize;
        $this->isPrime = array_fill(0, $n+1, false);
    }

    private function hydrateArray()
    {
        $n = $this->sieveSize;
        $this->isPrime[2] = true;
        for ($i = 3; $i <= $n+1; $i = $i + 2) {
            if ($i%2 != 0) {
                $this->isPrime[$i] = true;
            }
        }
    }

    public function runSieve()
    {
        $factor = 3;
        $q = sqrt($this->sieveSize);

        while ($factor < $q)
        {
            for ($num = $factor; $num < $this->sieveSize; $num = $num+2)
            {
                if ($this->isPrime[$num])
                {
                    $factor = $num;
                    break;
                }
            }
            $this->clearBit($factor);    
            $factor += 2;
        }
    }

    private function clearBit($factor)
    {
        for ($num = $factor * 3; $num < $this->sieveSize; $num += $factor * 2) {
            $this->isPrime[$num] = false;
        }
    }

    public function printResults(bool $showResults, float $duration, int $passes)
    {
        if ($showResults)
            printf("2, ");

        $count = 1;
        for ($num = 3; $num <= $this->sieveSize; $num++)
        {
            if ($this->isPrime[$num])
            {
                if ($showResults)
                    printf("%d, ", $num);
                $count++;
            }
        }

        if ($showResults)
            printf("\n");
        
        printf("Passes: %d, Time: %lf, Avg: %lf, Limit: %d, Count: %d, Valid: %d\n", 
               $passes, 
               $duration, 
               $duration / $passes, 
               $this->sieveSize, 
               $count, 
               $this->validateResults());
    }

    public function countPrimes(): int
    {
        $count = 0;
        for ($i = 0; $i < $this->sieveSize; $i++)
        if ($this->isPrime[$i])
                $count++;
        return $count;
    }    
}

$tstart = microtime(true);
$passes = 0;

while((microtime(true) - $tstart) < 10.00) {
    $sieve = new PrimeSieve(1000000);
    $sieve->runSieve();
    $passes++;
}

$tD = microtime(true) - $tstart;
if ($sieve) {
    $sieve->printResults(false, $tD, $passes);
}