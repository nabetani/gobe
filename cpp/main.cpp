#include <cstdio>
#include <cstdlib>
#include <cstdint>
#include <chrono>
#include <vector>
#include <functional>

namespace ch = std::chrono;
using cl = std::chrono::steady_clock;

uint32_t test(uint32_t num){
    std::vector<std::function<uint32_t(uint32_t)>> m;
    m.reserve(num);
    for( uint32_t i=0; i<num ; ++i ){
        m.emplace_back( [i,num](uint32_t a)->uint32_t{
			uint32_t x = (i+a)*7 + 11;
			uint32_t y = (x+a)*13 + 15;
			uint32_t z = (y+i)*17 + 19;
			uint32_t w = (z+i+a)*23 + 29;
			uint32_t t = x ^ (x << 11);
			return ((w ^ (w >> 19)) ^ (t ^ (t >> 8))) % num;
        });
    }
    uint32_t sum=0;
    for(uint32_t i=0 ; i<num ; i++){
        sum += [num,&m,i](uint32_t s)->uint32_t{
            std::vector<bool> b(num);
            for(;;){
                if (b[s]){
                    return s;
                }
                b[s]=true;
                s = m[s](i);
            }
        }(i);
        sum %= (1u<<24);
    }
    return sum;
}

int main( int argc, char const * argv[]){
    uint32_t num = argc<=1 ? 1000 : std::atoi(argv[1]);
    auto t0 = cl::now();
    uint32_t r = test(num);
    auto t1 = cl::now();
    auto ms = ch::duration_cast<ch::microseconds>(t1-t0).count() * 1e-3;
    printf( "r:%d, compiler:%s, tick:%.2fms\n" ,r, __VERSION__, ms);
    return 0;
}