#include <string>
#include <vector>
#include <iostream>

using namespace std;

vector<int> diStringMatch(string s)
{
    vector<int> arr;
    int i = 0;
    int k = 0;
    int len = s.length();
      int j = len;
    while (k < len)
    {
        if (s[k] == 'I')
            arr.push_back(i++);
        else
            arr.push_back(j--);
        k++;
    }
    if (s[k] == 'I')
            arr.push_back(i++);
        else
            arr.push_back(j--);
    return (arr);
}
int main()
{
    string a = "IDIDI";
    vector <int> arr = diStringMatch(a);
    for (int i = 0 ;i < arr.size(); i++)
    {
        cout << arr[i] << endl;
    }
}