#include <vector>
#include <map>
#include <string>
#include <algorithm>

using namespace std;

class Solution {
public:
    int longestCommonSubsequence(string text1, string text2) {
        int n = text1.length();
        int m = text2.length();

        int matrix[n+1][m+1];

        for (int i = 0; i < n + 1; i++) {
            for (int j = 0 ; j < m + 1 ;j++) {
                if (i == 0 || j == 0)
                    matrix[i][j] = 0;
                else if (text1[n - i] == text2[m - j]) {
                    matrix[i][j] = 1 + matrix[i - 1][j - 1];
                } else {
                    matrix[i][j] = max(matrix[i - 1][j], matrix[i][j - 1]);
                }
            }
        }
        return (matrix[n][m]);
    }
};

int main() {
    Solution sol;
}
