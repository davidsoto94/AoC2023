using Part1;

namespace Day5.Test;

public class UnitTest1
{
    [Fact]
    public void Test1()
    {
        var part1 = new Part1Class();


        var result = part1.Part1Function("../../../..");

        Assert.Equal(35, result);

    }
}