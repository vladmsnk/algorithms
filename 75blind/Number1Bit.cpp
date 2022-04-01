#include <vector>
#include <map>
#include <string>

using namespace std;


class Solution {
public:
    int hammingWeight(uint32_t n) {
        uint32_t cnt = 0;
        while (n > 0)
        {
            cnt++;
            n = n & (n - 1);
        }
        return cnt;
    }
};

int main() {
    Solution sol;
}