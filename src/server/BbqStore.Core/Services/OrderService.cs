using System;
using System.Collections.Generic;
using System.Text;
using BbqStore.Core.Entities;
using Marten;

namespace BbqStore.Core.Services
{
    public interface IOrderService : IEntityService<Order>
    {

    }
    public class OrderService : EntityService<Order>, IOrderService
    {
        public OrderService(IDocumentSession documentSession) : base(documentSession)
        {
        }


    }
}
