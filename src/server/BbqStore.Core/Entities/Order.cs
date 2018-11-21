using System.Collections.Generic;

namespace BbqStore.Core.Entities
{
    public class Order : Entity
    {
        public string CustomerName { get; set; }
        public string Status { get; set; }
        public List<OrderLine> Lines { get; set; }
    }
}