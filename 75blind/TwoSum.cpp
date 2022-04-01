#include <vector>
#include <map>
#include <string>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> indeces;
        map<int, int> check;

        for (int i = 0; i < nums.size(); i++) {
            if (check.find(target - nums[i]) != check.end())
            {
                indeces.push_back(check[target - nums[i]]);
                indeces.push_back(i);
                break;
            }
            check[nums[i]] = i;
        }
        return indeces;
    }
};

int main() {
    Solution sol;
}