using System;

namespace BbqStore.Core.Entities
{
    public class OrderLine : Entity
    {
        public int Quantity { get; set; }
        public Guid ProductId { get; set; }
        public decimal LineTotal { get; set; }
    }

    public class DisplayOrderLine : OrderLine
    {
        public DisplayOrderLine(OrderLine orderLine)
        {
            Id = orderLine.Id;
            Quantity = orderLine.Quantity;
            ProductId = orderLine.ProductId;
            LineTotal = orderLine.LineTotal;
            CreatedBy = orderLine.CreatedBy;
            CreatedDate = orderLine.CreatedDate;
            ModifiedBy = orderLine.ModifiedBy;
            ModifiedDate = orderLine.ModifiedDate;
        }
        public string ProductName { get; set; }
        public string ProductUnit { get; set; }
    }
}