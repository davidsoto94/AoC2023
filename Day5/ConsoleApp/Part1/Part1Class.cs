using Microsoft.VisualBasic;

namespace Part1;
public class Part1Class
{
    Dictionary<long, long> SeedToSoil = new Dictionary<long, long>();
    Dictionary<long, long> SoilToFertilizer = new Dictionary<long, long>();
    Dictionary<long, long> FertilizerToWater = new Dictionary<long, long>();
    Dictionary<long, long> WaterToLight = new Dictionary<long, long>();
    Dictionary<long, long> LightToTemperature = new Dictionary<long, long>();
    Dictionary<long, long> TemperatureToHumidity = new Dictionary<long, long>();
    Dictionary<long, long> HumidityToLocation = new Dictionary<long, long>();
    public long Part1Function(string RootPath){
        IEnumerable<string> lines = File.ReadLines(RootPath+"/Files/Test.txt");

        var Seeds = new List<long>();
        
        Dictionary<long, long> selectedDictionary = null;
        foreach(string line in lines){
            if(line.Length==0){
                continue;
            }
            string[] part = line.Split(" ");
            if( part[0].Contains("seeds")){
                for (long i = 1; i < part.Length;i++){
                    //Console.WriteLine(part[i].Trim());
                    Seeds.Add(long.Parse(part[i].Trim()));
                }
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
        Console.WriteLine("test");
        long soil=1000000000;
        foreach( long seed in Seeds){
            long temp = GetSoil(seed);
            if (temp<soil){
                soil = temp;
            }
        }
        
        return soil;
    }


    public long GetSoil(long seed){
        long result = seed;
        if(SeedToSoil.ContainsKey(seed)){
            result = SeedToSoil[seed];
        }
        if(SoilToFertilizer.ContainsKey(result)){
            result = SoilToFertilizer[result];
        }
        if(FertilizerToWater.ContainsKey(result)){
            result = FertilizerToWater[result];
        }
        if(WaterToLight.ContainsKey(result)){
            result = WaterToLight[result];
        }
        if(LightToTemperature.ContainsKey(result)){
            result = LightToTemperature[result];
        }
        if(TemperatureToHumidity.ContainsKey(result)){
            result = TemperatureToHumidity[result];
        }
        if(HumidityToLocation.ContainsKey(result)){
            result = HumidityToLocation[result];
        }
        return result;
    }

    public void AddToDict(string[] line, Dictionary<long, long> Dict){
        for(long i = 0; i<long.Parse(line[2]);i++){
            Dict[long.Parse(line[1]) + i] = long.Parse(line[0]) + i;
        }
    }
}
