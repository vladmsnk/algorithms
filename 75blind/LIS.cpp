#include <vector>
#include <map>
#include <string>
#include <algorithm>


using namespace std;


class Solution {
public:

    int lengthOfLIS(vector<int>& nums) {
        int  j;
        vector<int> lengths(nums.size(), 1);

        for (j = 1; j < nums.size(); j++) { 
            for (int i = 0; i < j; i++ ) {
                if (nums[j] > nums[i]) {
                    lengths[j] = max(lengths[j], lengths[i] + 1);
                }
            }

        }
        auto it = max_element(lengths.begin(), lengths.end());
        return (*it);
    }
};

int main() {
    Solution sol;
}