using System.Text;
using Microsoft.Extensions.Configuration;

namespace BbqStore.Core.Services
{
    public class ConfigurationService : IConfigurationService
    {
        public ConfigurationService(IConfiguration config)
        {
            Configuration = config;
        }

        protected IConfiguration Configuration { get; set; }

        public string Get(string key)
        {
            return Configuration[key];
        }
    }
}
