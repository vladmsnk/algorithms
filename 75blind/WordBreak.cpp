#include <vector>
#include <map>
#include <string>

using namespace std;


class Solution {
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        int len = s.length();
        bool    check[len + 1];
        for (int i = 0;i < len + 1; i++) {
            check[i] = false;
        }
        int i = len - 1;
        check[len] = true;
        while (i != -1) {
            string str = s.substr(i, len - i);
            for (auto word : wordDict) {
                if (str.find(word) == 0) {
                    if (check[i] == false)
                        check[i] = check[i + word.length()];
                    if (check[i] == true)
                        break;
                }
            }
            i--;
        }
        return check[0];
    }
};

int main() {
    Solution sol;
}