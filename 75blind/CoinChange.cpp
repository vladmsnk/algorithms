#include <vector>
#include <map>
#include <string>
#include <algorithm>

using namespace std;

class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        int check[amount + 1];
    
        check[0] = 0;
        for (int i = 1; i < amount + 1; i++) {
            check[i] = amount + 1;
        }

        for (int i = 1; i < amount + 1; i++) {
            for (int j = 0; j < coins.size(); j++) {
                if (i - coins[j] >= 0) {
                    check[i] = min(check[i - coins[j]] + 1, check[i]);
                }
            }
        }
        if (check[amount] == amount + 1)
            return -1;
        return (check[amount]);

    }
};


int main() {
    Solution sol;
}