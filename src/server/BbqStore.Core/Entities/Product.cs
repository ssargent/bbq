using System;
using System.Collections.Generic;
using System.Text;

namespace BbqStore.Core.Entities
{
    public class Entity
    {
        public virtual Guid Id { get; set; }
        public virtual string CreatedBy { get; set; }
        public virtual DateTimeOffset CreatedDate { get; set; }
        public virtual string ModifiedBy { get; set; }
        public virtual DateTimeOffset ModifiedDate { get; set; }
        public virtual bool IsDeleted { get; set; }
    }

    public class NamedEntity : Entity
    {
        public virtual string Key { get; set; }
        public virtual string Name { get; set; }
    }
    public class Product : NamedEntity
    {
        public string Description { get; set; }
        public Decimal Price { get; set; }
        public string Unit { get; set; }
    }

    public class Order : Entity
    {
        public string CustomerName { get; set; }
        public string Status { get; set; }
        public List<OrderLine> Lines { get; set; }
    }

    public class OrderLine : Entity
    {
        public int Quantity { get; set; }
        public Guid ProductId { get; set; }
        public decimal LineTotal { get; set; }
    }

    public class Store : NamedEntity
    {

    }
}
