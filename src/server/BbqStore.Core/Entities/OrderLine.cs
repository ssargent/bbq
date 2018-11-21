using System;

namespace BbqStore.Core.Entities
{
    public class OrderLine : Entity
    {
        public int Quantity { get; set; }
        public Guid ProductId { get; set; }
        public decimal LineTotal { get; set; }
    }
}