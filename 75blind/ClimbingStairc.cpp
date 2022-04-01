#include <vector>
#include <map>
#include <string>

using namespace std;


class Solution {
public:
    int climbStairs(int n) {
        int cnt[n];
        cnt[0] = 1;
        if (n == 1)
            return (1);
        cnt[1] = 2;
        for (int i = 2; i < n; i++) {
            cnt[i] = cnt[i- 1] + cnt[i - 2];
        }
        return (cnt[n - 1]);
    }
};

int main() {
    Solution sol;
}