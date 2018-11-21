using System.Collections.Generic;
using System.Text;
using Marten;
using Marten.Schema;

namespace BbqStore.Core.Database
{
    public class BbqStoreInitialData : IInitialData
    {
        private object[] InitialData { get; set; }

        public BbqStoreInitialData(params object[] initialData)
        {
            InitialData = initialData;
        }

        public void Populate(IDocumentStore store)
        {
            using (var session = store.LightweightSession())
            {
                session.Store(InitialData);
                session.SaveChanges();
            }
        }
    }
}