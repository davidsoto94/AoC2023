using System.Numerics;
using Microsoft.VisualBasic;

namespace Part1;
public class Part1Class
{   
    List<List<long>> SeedToSoil = new List<List<long>>();
    List<List<long>> SoilToFertilizer = new List<List<long>>();
    List<List<long>> FertilizerToWater = new List<List<long>>();
    List<List<long>> WaterToLight = new List<List<long>>();
    List<List<long>> LightToTemperature = new List<List<long>>();
    List<List<long>> TemperatureToHumidity = new List<List<long>>();
    List<List<long>> HumidityToLocation = new List<List<long>>();
    
    public long Part1Function(string RootPath){
        IEnumerable<string> lines = File.ReadLines(RootPath+"/Files/Test.txt");

        var Seeds = new List<long>();
        
        string[] part = lines.ToArray()[0].Split(" ");
        if( part[0].Contains("seeds")){
            for (long i = 1; i < part.Length;i++){
                //Console.WriteLine(part[i].Trim());
                Seeds.Add(long.Parse(part[i].Trim()));
            }
        }

        List<List<long>> selectedDictionary = null;

        foreach(string line in lines){
            if(line.Length==0){
                continue;
            }
            part = line.Split(" ");
            if( part[0].Contains("seeds")){
                continue;
            }
            if (!long.TryParse(part[0].Trim(),out long val)){
                switch (part[0]) {
                    case "seed-to-soil":
                        selectedDictionary = SeedToSoil;
                        break;
                    case "soil-to-fertilizer":
                        selectedDictionary = SoilToFertilizer;
                        break;
                    case "fertilizer-to-water":
                        selectedDictionary = FertilizerToWater;
                        break;
                    case "water-to-light":
                        selectedDictionary = WaterToLight;
                        break;
                    case "light-to-temperature":
                        selectedDictionary = LightToTemperature;
                        break;
                    case "temperature-to-humidity":
                        selectedDictionary = TemperatureToHumidity;
                        break;
                    case "humidity-to-location":
                        selectedDictionary = HumidityToLocation;
                        break;
                    
                }
                continue;
            }
            AddToDict(part, selectedDictionary);


            //Console.WriteLine(selectedDictionary);
        }
        //Console.WriteLine(string.Join(" ", FertilizerToWater.SelectMany(innerList => innerList)));
        long soil=1000000000000000000;
        foreach( long seed in Seeds){
            long result = seed;
            result = GetMappedValue(result, SeedToSoil);
            result = GetMappedValue(result, SoilToFertilizer);
            result = GetMappedValue(result, FertilizerToWater);
            result = GetMappedValue(result, WaterToLight);
            result = GetMappedValue(result, LightToTemperature);
            result = GetMappedValue(result, TemperatureToHumidity);
            result = GetMappedValue(result, HumidityToLocation);
            if(result<soil){
                soil = result;
            }
            
        }
        return soil;
    }

    public long GetMappedValue(long InsertValue, List<List<long>> Map){
        long tempValue = InsertValue;
        foreach (List<long> line in Map)
        {
            if (InsertValue >= line[1] && InsertValue < line[1]+line[2]){
                tempValue = line[0] + (InsertValue-line[1]);
            }
        }
        return tempValue;
    }

    public void AddToDict(string[] line, List<List<long>> list){
        list.Add(new List<long>(){
            long.Parse(line[0]),long.Parse(line[1]),long.Parse(line[2])
        });
        
    }
}
