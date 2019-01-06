using System.Collections.Generic;
using System.Linq;

namespace BbqStore.Core.Entities
{
    public class Order : Entity
    {
        public string CustomerName { get; set; } = "anonymous";
        public string Status { get; set; } = "draft";
        public List<OrderLine> Lines { get; set; } = new List<OrderLine>();
        public decimal Total => Lines.Sum(l => l.LineTotal);
    }

    public class DisplayOrder : Entity
    {
        public string CustomerName { get; set; } = "anonymous";
        public string Status { get; set; } = "draft";
        public List<DisplayOrderLine> Lines { get; set; } = new List<DisplayOrderLine>();
        public decimal Total => Lines.Sum(l => l.LineTotal);
    }

}