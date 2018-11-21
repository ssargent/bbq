using System;
using System.Text;

namespace BbqStore.Core.Entities
{
    public class Product : NamedEntity
    {
        public string Description { get; set; }
        public Decimal Price { get; set; }
        public string Unit { get; set; }
    }
}
