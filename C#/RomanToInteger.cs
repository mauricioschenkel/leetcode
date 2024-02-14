public class Solution {
    public int RomanToInt(string s) {
        
        int convertedInt = 0;
        Dictionary<char, int> romanNums = new Dictionary<char, int>();
        romanNums['I'] = 1;
        romanNums['V'] = 5;
        romanNums['X'] = 10;
        romanNums['L'] = 50;
        romanNums['C'] = 100;
        romanNums['D'] = 500;
        romanNums['M'] = 1000;

        for (int i = 0; i < s.Length; i++)
        {
            char curr = s[i];
            char next = i < (s.Length - 1) ? s[i + 1] : s[i];

            if (romanNums[curr] >= romanNums[next])
                convertedInt += romanNums[curr];
            else
            {
                convertedInt += (romanNums[next] - romanNums[curr]);
                i++;
            }
        }

        return convertedInt;
    }
}