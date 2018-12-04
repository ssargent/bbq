using System;
using System.Linq;
using System.Collections.Generic;
using System.Text;
using BbqStore.Core.Entities;
using Marten;

namespace BbqStore.Core.Services
{
    public interface IOrderService : IEntityService<Order>
    {
        Order AddItemToOrder(Guid id, OrderLine line);
    }
    public class OrderService : EntityService<Order>, IOrderService
    {
        public OrderService(IDocumentSession documentSession) : base(documentSession)
        {
        }

        public Order AddItemToOrder(Guid id, OrderLine line)
        {
            var order = default(Order);

            order = Guid.Empty.Equals(id) ? Save(new Order()) : GetById(id);

            var product = DocumentSession.Query<Product>().FirstOrDefault(p => p.Id == line.ProductId);

            if (order.Lines.Any(l => l.ProductId == line.ProductId))
            {
                var existingLine = order.Lines.FirstOrDefault(ol => ol.ProductId == line.ProductId);
                existingLine.Quantity += 1;
                existingLine.LineTotal = product.Price * existingLine.Quantity;
                existingLine.ModifiedDate = DateTimeOffset.Now;
            }
            else
            {

                line.LineTotal = line.Quantity * product.Price;

                line.CreatedBy = "chef";
                line.CreatedDate = DateTimeOffset.Now;
                line.ModifiedBy = "chef";
                line.ModifiedDate = DateTimeOffset.Now;
                order.Lines.Add(line);
            }

            order.ModifiedBy = "chef";
            order.ModifiedDate = DateTimeOffset.Now;

            Save(order);

            return order;
        }
    }
}
