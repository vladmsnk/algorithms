#include <vector>
#include <map>
#include <string>
#include <algorithm>

using namespace std;

//O(n^2) needs thinking for O(nlogn)
class Solution {
public:
    vector<int> countBits(int n) {
        vector<int> ans;
    
        for (int i = 0; i < n + 1; i++) {
            int cnt = 0;
            int copy_i = i;
            while (copy_i > 0) {
                cnt++;
                copy_i = copy_i & (copy_i - 1);
            }
            ans.push_back(cnt);
        }
        return ans;
    }
};


int main() {
    Solution sol;
}