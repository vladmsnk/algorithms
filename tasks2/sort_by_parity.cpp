#include <vector>
#include <iostream>

using namespace std;

vector<int> sortArrayByParity(vector<int>& nums) 
    {
        int i = 0;
        int j = nums.size() - 1;
        while (i < j)
        {
            if (nums[j] % 2 == 0)
            {
                int tmp = nums[i];
                nums[i] = nums[j];
                nums[j] = tmp;
                i++;
            }
            else
                j--;
        }
        return (nums);
    }
int main()
{
    vector<int> arr = {3,2};
    vector<int> sorted = sortArrayByParity(arr);
    for (int i = 0 ; i < sorted.size(); i++)
    {
        cout << sorted[i] << endl;
    }
}